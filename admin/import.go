package admin

import (
	"bufio"
	"io"
	"net/url"
	"regexp"
	"strings"

	"github.com/crhym3/wrabble-game/wrabble"
	"github.com/crhym3/wrabble-game/wrabble/ds"

	"appengine"
	"appengine/blobstore"
	"appengine/datastore"
	"appengine/taskqueue"
)

var reWordImport = regexp.MustCompile(`^[a-z]{3,}$`)

func asyncImportWords(c appengine.Context, dict string, words []string) error {
	t := taskqueue.NewPOSTTask(pathImportWords, url.Values{
		"dict":  {dict},
		"words": words,
	})
	_, err := taskqueue.Add(c, t, "")
	return err
}

func asyncImportDict(c appengine.Context, b appengine.BlobKey, name string) error {
	t := taskqueue.NewPOSTTask(pathImportDict, url.Values{
		"blobkey": {string(b)},
		"name":    {name},
	})
	_, err := taskqueue.Add(c, t, "")
	return err
}

func importWords(c appengine.Context, dict string, words []string) error {
	if len(words) == 0 {
		return nil
	}
	c.Infof("importing %d words starting with %q", len(words), words[0])
	keys := make([]*datastore.Key, 0, len(words))
	items := make([]*wrabble.Word, 0, len(words))
	for _, word := range words {
		k, w := ds.NewWord(c, dict, word)
		keys = append(keys, k)
		items = append(items, w)
	}
	_, err := datastore.PutMulti(c, keys, items)
	return err
}

func importDict(c appengine.Context, b appengine.BlobKey, dict string) error {
	c.Infof("importing dict from %q", b)
	defer blobstore.Delete(c, b)
	r := bufio.NewReader(blobstore.NewReader(c, b))

	var total, skipped int
	words := []string{}

	for {
		line, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			c.Errorf("error reading from blob: %v", err)
			return err
		}

		line = strings.Trim(line, " \n\r")

		if !reWordImport.Match([]byte(line)) {
			skipped++
			if err == io.EOF {
				break
			}
			continue
		}

		total++
		words = append(words, line)
		if len(words) > 99 {
			asyncImportWords(c, dict, words)
			words = words[:0]
		}

		if err == io.EOF {
			break
		}
	}

	if len(words) > 0 {
		asyncImportWords(c, dict, words)
	}

	c.Infof("importDict: %d total, %d skipped", total, skipped)
	return nil
}
