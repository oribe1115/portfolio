var store = {
  debug: true,
  state: {
    target: "about",
    contentList: true,
    categoryID: "",
    contentID: ""
  },
  setToMainCategory(id) {
    this.state.target = "mainCategory";
    this.state.categoryID = id;
  },
  setToSubCategory(id) {
    this.state.target = "subCategory";
    this.state.categoryID = id;
  },
  setToAbout() {
    this.state.target = "about";
    this.state.categoryID = "";
  }
};

export default store;
