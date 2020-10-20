package menu

type MenuItems struct {
	ItemName,
	ItemIcon,
	ItemLink string
}

var defaultMenu []MenuItems = []MenuItems{
	MenuItems{
		ItemName: "Dashboard",
		ItemIcon: "dashboard",
		ItemLink: "/member/dashboard",
	},
	MenuItems{
		ItemName: "My Files",
		ItemIcon: "cloud",
		ItemLink: "/member/myfiles",
	},
}

func GetMenu() []MenuItems {
	return defaultMenu
}
