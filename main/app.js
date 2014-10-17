function bootstrapApp(t) {
  t.menuItemSelected = function(e, detail, sender) {
    if (detail.isSelected) {
      this.$ && this.$.scaffold.closeDrawer();
    }
  };
}
