import { StyleSheet, Dimensions } from "react-native";

const sidebarWidth = Dimensions.get("window").width * 0.75;

const styles = StyleSheet.create({
  container: {
    flex: 1,
   
  },
  header: {
    flexDirection: "row",
    paddingTop: 16,
    paddingHorizontal: 10,
    alignItems: "center",
    justifyContent: "space-between",
  },
  hamburgerIcon: {
    width: 32,
    height: 32,
  },
  searchContainer: {
    flex: 1,
    marginLeft: 10,
  },
  overlay: {
    position: "absolute",
    top: 0,
    left: 0,
    bottom: 0,
    right: 0,
    backgroundColor: "rgba(0,0,0,0.5)",
    justifyContent: "flex-start",
    zIndex: 1,
  },
  sidebar: {
    position: "absolute",
    top: 0,
    left: 0,
    height: "100%",
    backgroundColor: "#fff",
    padding: 20,
    shadowColor: "#000",
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.8,
    shadowRadius: 2,
    elevation: 5,
    zIndex: 2,
    width: sidebarWidth,
  },
  sidebarTitle: {
    fontSize: 24,
    fontWeight: "bold",
    marginBottom: 20,
  },
  menuItem: {
    paddingVertical: 15,
    borderBottomWidth: 1,
    borderBottomColor: "#ccc",
  },
  menuItemText: {
    fontSize: 18,
  },
  contentContainer: {
    flex: 1,
    paddingTop: 10,
  },
  flatListContainer: {
    alignItems: "center",
  },
  cardContainer: {
    margin: 12,
  },
  floatingButton: {
    position: "absolute",
    bottom: 24,
    left: 0,
    right: 0,
    justifyContent: "center",
    alignItems: "center",
  },
  plusIcon: {
    width: 56,
    height: 56,
  },
});

export default styles;
