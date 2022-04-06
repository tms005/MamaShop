package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var page int = 0
var category = []string{"Household", "Food", "Drinks"}
var itemsDetail = map[string]Item{
	"Cups":   {category: 0, unit: 5, cost: 3},
	"Cake":   {category: 1, unit: 3, cost: 1},
	"Sprite": {category: 2, unit: 5, cost: 2},
	"Fork":   {category: 0, unit: 4, cost: 3},
	"Bread":  {category: 1, unit: 2, cost: 2},
	"Plates": {category: 0, unit: 4, cost: 3},
	"Coke":   {category: 2, unit: 5, cost: 2},
}
var shopitemslist []ShoppingList
var items = setItems(itemsDetail) // store key of map
func setItems(mp map[string]Item) (item []string) {
	for i, _ := range mp {
		item = append(item, i)
	}
	return
}

type Item struct {
	category int
	unit     int
	cost     float64
}

type ShoppingList struct {
	ItemName string
	Quantity int
}

func main() {
	for {
		pages(page)
	}
}

// 1st code to run
func pages(page int) {
	fmt.Println()
	var s int
	switch page {
	case 0:
		s, _ = fmt.Println("Shopping List Application")
		fmt.Println(strings.Repeat("=", s-1))
		menu()
	case 1:
		s, _ = fmt.Println("Shopping List Contents")
		fmt.Println(strings.Repeat("=", s-1))
		shoppingListContents()
	case 2:
		s, _ = fmt.Println("Generate Report")
		fmt.Println(strings.Repeat("=", s-1))
		generateSLReport()
	case 3:
		s, _ = fmt.Println("Add Items")
		fmt.Println(strings.Repeat("=", s-1))
		addItems()
	case 4:
		s, _ = fmt.Println("Modify Items")
		fmt.Println(strings.Repeat("=", s-1))
		modifyItem()
	case 5:
		s, _ = fmt.Println("Delete Items")
		fmt.Println(strings.Repeat("=", s-1))
		deleteItem()
	case 6:
		s, _ = fmt.Println("Print Current Data")
		fmt.Println(strings.Repeat("=", s-1))
		printCurData()
	case 7:
		s, _ = fmt.Println("Add New Category Name")
		fmt.Println(strings.Repeat("=", s-1))
		addCategery()
	case 8:
		s, _ = fmt.Println("Modify Category Name")
		fmt.Println(strings.Repeat("=", s-1))
		modifyCategery()
	case 9:
		s, _ = fmt.Println("Remove Category Name")
		fmt.Println(strings.Repeat("=", s-1))
		removeCategery()
	case 10:
		s, _ = fmt.Println("Shopping List")
		fmt.Println(strings.Repeat("=", s-1))
		shoppingList()
	default:
		fmt.Println("Invalid input!")
		pages(0)
	}
}

//0. Main Menu
func menu() {
	fmt.Println("1. View entire shopping list")
	fmt.Println("2. Generate Shopping List Report")
	fmt.Println("3. Add Items")
	fmt.Println("4. Modify Items")
	fmt.Println("5. Delete Item")
	fmt.Println("6. Print Current Data.")
	fmt.Println("7. Add New Category Name")
	fmt.Println("8. Modify Category Name")
	fmt.Println("9. Remove Category Name")
	fmt.Println("10. Shopping List")
	fmt.Println("Select Your Choice:")
	input := userInput()
	v, err := strconv.Atoi(input)
	if len(input) == 0 {
		fmt.Println("No Input Found!")
	} else if err != nil {
		fmt.Println("Invalid input!")
	} else {
		page = v
	}
	return
}

//1. Shopping List Contents
func shoppingListContents() {
	if len(itemsDetail) == 0 {
		fmt.Println("No item recoded!")
		page = 0
		userInput()
		return
	}
	for _, v := range items {
		printFormat(v, itemsDetail[v])
	}
	page = 0
	userInput()
}

//2. Generate Shopping List Report
func generateSLReport() {
	if len(itemsDetail) == 0 {
		fmt.Println("No item recoded!")
		page = 0
		userInput()
		return
	}
	fmt.Println("1. Total cost of each category.")
	fmt.Println("2. List of item by category.")
	fmt.Println("3. Main Menu.")
	fmt.Println("Choose your report:")
	input := userInput()
	fmt.Println()
	if input == "1" {
		s, _ := fmt.Println("Total cost by Category.")
		fmt.Println(strings.Repeat("-", s-1))
		for i, v := range category {
			fmt.Printf("%v cost : %v\n", v, totalUpValue(i))
		}
		fmt.Println("Continue press enter:")
		userInput()
	} else if input == "2" {
		s, _ := fmt.Println("List by Category.")
		fmt.Println(strings.Repeat("-", s-1))
		for i, _ := range category {
			for k, v := range itemsDetail { // notes: if without i the printout not in order
				if v.category == i {
					printFormat(k, v)
				}
			}
		}
		fmt.Println("Continue press enter:")
		userInput()
	} else if input == "3" || input == "" {
		page = 0
	} else {
		fmt.Println("Invalid input!")
	}

}

//3. Add Items
func addItems() {

	var (
		input string
		inNa  string
		inCa  string
		inCaI int
		inUn  int
		inCo  float64
		err   error
	)
	//Name
	if len(category) == 0 {
		fmt.Println("There is not category name defind yet! Do you want to go to Add New Category?: (Key 'Yes' to go)")
		input = userInput()
		if strings.ToUpper(input) == "YES" || strings.ToUpper(input) == "Y" {
			page = 7
			return
		} else {
			page = 0
			return
		}
	} else {
		fmt.Println("What is the name of your item? (Enter to main menu)") //Enter to main menu
		inNa = userInput()
		if inNa == "" {
			page = 0
			return
		}
	}
	//Check Repeated Item
	if len(items) > 0 {
		for i := 0; i < len(items); i++ { //Come back 'to update to pull out function'
			if inNa == (items[i]) {
				fmt.Printf("Item: %s is already exist! (Suggest to use different item's name) \n", inNa)
				userInput()
				return
			}
		}
	}
	//Category
	for {
		fmt.Println("What category does it belong to?")
		inCa = userInput()
		check, i, _ := checkData(inCa)
		inCaI = i
		if inCa == "" {
			fmt.Println("No Input Found!")
			continue
		} else if check == true { //Check Name exist in category
			break
		} else {
			fmt.Printf("%s is not add in category yet! Do you want to go to Add New Category?: (Key 'Yes' to go or enter to continue)\n", inCa)
			input = userInput()
			if strings.ToUpper(input) == "YES" || strings.ToUpper(input) == "Y" {
				page = 7
				return
			}
			continue
		}
	}
	//Units
	for {
		fmt.Println("How many units are there?")
		inUn, err = strconv.Atoi(userInput())
		if err != nil {
			fmt.Println("Invalid input!")
			continue
		} else {
			break
		}

	}
	//Cost
	for {
		fmt.Println("How much does it cost per unit?")
		inCo, err = strconv.ParseFloat(userInput(), 10)
		if err != nil {
			fmt.Println("Invalid input!")
			continue
		} else {
			break
		}
	}

	//add to struct
	temp := Item{
		category: inCaI,
		unit:     inUn,
		cost:     inCo,
	}
	items = append(items, inNa)
	itemsDetail[inNa] = temp
	// fmt.Println(itemsDetail) //Print Check
	// fmt.Printf("%#v", items) //test Print Check
}

//4. Modify Items
func modifyItem() {

	var tIndex int
	var found bool
	fmt.Println("Which item would you wish to modify? (Enter to main menu)")
	input := userInput()

	if input == "" {
		page = 0
		return
	}

	for i, v := range items {
		if input == v {
			tIndex = i
			found = true
			break
		}
	}

	if found != true {
		fmt.Println("Item name not found")
		return
	}

	fmt.Printf("Current item name is %v - Category is %v - Quantity is %v - Unit Cost %v\n", items[tIndex], category[itemsDetail[items[tIndex]].category], itemsDetail[items[tIndex]].unit, itemsDetail[items[tIndex]].cost)
	type tempRecord struct {
		change bool
		reply  string
		data   Item
	}
	var ipdata []tempRecord

	//Change Item's Name
	fmt.Println("Enter new Name. Enter for no change.")
	input1 := userInput()
	c, t, da := checkInputModifyValue(1, tIndex, input1)
	ip1 := tempRecord{
		change: c,
		reply:  t,
		data:   da,
	}
	ipdata = append(ipdata, ip1)

	//Change Category's Name
	fmt.Println("Enter new Category. Enter for no change.")
	input2 := userInput()
	c, t, da = checkInputModifyValue(2, tIndex, input2)
	ip2 := tempRecord{
		change: c,
		reply:  t,
		data:   da,
	}
	ipdata = append(ipdata, ip2)

	//Change Quantity
	fmt.Println("Enter new Quantity. Enter for no change.")
	input3 := userInput()
	c, t, da = checkInputModifyValue(3, tIndex, input3)
	ip3 := tempRecord{
		change: c,
		reply:  t,
		data:   da,
	}
	ipdata = append(ipdata, ip3)

	fmt.Println("Enter new Unit Cost. Enter for no change.")
	input4 := userInput()
	c, t, da = checkInputModifyValue(4, tIndex, input4)
	ip4 := tempRecord{
		change: c,
		reply:  t,
		data:   da,
	}
	ipdata = append(ipdata, ip4)
	// fmt.Println(ipdata) //Print Test

	temp := Item{}
	titemsDetail := itemsDetail[items[tIndex]]
	titems := items[tIndex]
	for i, v := range ipdata {
		// fmt.Printf("record no%v. %v - %T   %v\n", i, v.data, v.data, v.change)
		switch i {
		case 0:
			if v.change {
				items[tIndex] = v.reply
			} else {
				fmt.Printf("%v\n", v.reply)
			}
		case 1:
			if v.change {
				temp.category = v.data.category
			} else {
				temp.category = titemsDetail.category
				fmt.Printf("%v\n", v.reply)
			}
		case 2:
			if v.change {
				temp.unit = v.data.unit
			} else {
				temp.unit = titemsDetail.unit
				fmt.Printf("%v\n", v.reply)
			}
		case 3:
			if v.change {
				temp.cost = v.data.cost
			} else {
				temp.cost = titemsDetail.cost
				fmt.Printf("%v\n", v.reply)
			}
		}
	}
	delete(itemsDetail, titems)
	itemsDetail[items[tIndex]] = temp //Data Updated Here
}

//5. Delete Item
func deleteItem() {
	if len(items) == 0 {
		fmt.Println("No item recoded!")
		page = 0
		return
	}
	fmt.Println("Enter item name to delete: (Enter to main menu)")
	input := userInput()
	if len(input) == 0 {
		page = 0
		return
	} else { //revise it to be easy read******
		var existName bool
		var tIndex int
		for i := 0; i < len(items); i++ {
			if items[i] == input {
				tIndex = i
				existName = true
				break
			}
		}
		if !existName {
			fmt.Println("Item not found. Nothing to delete!")
			return
		}
		fmt.Printf("Confirm to delete %v with the data %v? (Enter 'Yes' to confirm)\n", items[tIndex], itemsDetail[items[tIndex]])
		input3 := userInput()
		if strings.ToUpper(input3) == "YES" || strings.ToUpper(input3) == "Y" {
			delete(itemsDetail, items[tIndex])
			items = append(items[:tIndex], items[tIndex+1:]...) //delete slice code
			fmt.Printf("Deleted %s!\n", input)
		} else {
			fmt.Println("Invalid input! Nothing to delete!")
		}
	}
	return
}

//6. Print Current Data.
func printCurData() {
	if len(items) <= 0 {
		fmt.Println("no data found! (Enter to continue)")
	} else {
		for _, v := range items {
			fmt.Printf("%v\t\t- %v\n", v, itemsDetail[v])
		}
	}
	page = 0
	userInput()
}

//7. add Category Name.
func addCategery() {
	fmt.Println("What is the New Category Name to add? (Enter to return main menu)")
	input := userInput()

	// pull out function
	condition := funcCategory(input)
	if condition == false {
		return
	} else {
		category = append(category, input)
		fmt.Printf("New category: %s added at index %v\n", input, len(category)-1)
		return
	}
	// fmt.Println(category) //Test result
}

//8. Modify Category Name
func modifyCategery() {
	var tIndex int
	if len(category) == 0 {
		fmt.Println("No item recoded!")
		page = 0
		userInput()
		return
	}

	for i, v := range category {
		fmt.Printf("%v. %v\n", i+1, v)
	}

	for {
		fmt.Println("Enter No in the category to edit: (Or enter to main menu)")
		input := userInput()
		input1, err := strconv.Atoi(input)
		if input == "" {
			page = 0
			return
		} else if err != nil {
			fmt.Println("Invalid input!")
			continue
		} else if input1 <= 0 || input1 > len(category) {
			fmt.Println("Invalid input!")
			continue
		}
		tIndex = input1 - 1
		break
	}
	for {
		fmt.Println("Input new category name: (Or enter * to main menu)")
		input := userInput()
		if input == "*" {
			page = 0
			return
		} else if input == "" {
			fmt.Println("Invalid input!")
		}

		condition := funcCategory(input)
		if condition == false {
			continue
		} else {
			tName := category[tIndex]
			category[tIndex] = input
			fmt.Printf("Category: change from %v to %v.  Enter to continue...\n", tName, category[tIndex])
			userInput()
			return
		}
	}

}

//9. Remove Category Name
func removeCategery() {
	var tIndex int
	var tName string
	if len(category) == 0 {
		fmt.Println("No item recoded!")
		page = 0
		userInput()
		return
	}

	for i, v := range category {
		fmt.Printf("%v. %v\n", i+1, v)
	}

	for {
		fmt.Println("Enter no of category to delete: (Enter to main menu)")
		input, err := strconv.Atoi(userInput())
		if input == 0 {
			page = 0
			return
		} else if err != nil {
			fmt.Println("Invalid input!")
			continue
		} else if input <= 0 || input > len(category) {
			fmt.Println("Invalid input!")
			continue
		} else {
			fmt.Printf("To delete %v in category, All RELATED DATA WILL BE DELETED!! (To confirm enter 'Yes')\n", category[input-1])
			input2 := userInput()
			if strings.ToUpper(input2) == "YES" {
				tIndex = input - 1
				break
			} else if strings.ToUpper(input2) == "NO" || strings.ToUpper(input2) == "N" {
				return
			}
			fmt.Println("Invalid input!")
			return
		}
	}
	var tStringMap []string
	for i := len(items) - 1; i >= 0; i-- {
		if itemsDetail[items[i]].category == tIndex {
			tStringMap = append(tStringMap, items[i])
			items = append(items[:i], items[i+1:]...)
		}
	}
	for _, v := range tStringMap {
		delete(itemsDetail, v)
	}
	for k, v := range itemsDetail {
		if v.category >= tIndex {
			temp := itemsDetail[k]
			temp.category = v.category - 1
			itemsDetail[k] = temp
		}
	}
	tName = category[tIndex]
	category = append(category[:tIndex], category[tIndex+1:]...)
	fmt.Printf("Category: %v and all related data deleted!!\nEnter to continue...", tName)
	userInput()
	page = 0
	return
}

//10. Shopping List
func shoppingList() {
	var nextI int
	for {
		fmt.Println("1. save shopping list")
		fmt.Println("2. retrieve previous shopping list")
		fmt.Println("Choose no: (Or enter to main menu)")
		SSLinput := userInput()
		fmt.Println()
		SSLinput1, err := strconv.Atoi(SSLinput)
		if SSLinput == "" {
			page = 0
			return
		} else if err != nil {
			fmt.Println("Invalid input!")
			return
		} else if SSLinput1 == 1 {
			nextI = 1
			break
		} else if SSLinput1 == 2 {
			nextI = 2
			break
		} else {
			fmt.Println("Invalid input!")
			return
		}
	}

	if nextI == 1 {
		b := saveShoppingList()
		if b == false {
			return
		}
	} else if nextI == 2 {
		b := retrieveShoppingList()
		if b == false {
			return
		}
	}

}

//Option 1 - Save Shopping List
func saveShoppingList() bool {
	var tInt int
	var tShopItemName []string
	for {
		s, _ := fmt.Println("Save Shopping List")
		fmt.Println(strings.Repeat("-", s-1))
		for i, s := range category {
			fmt.Printf("%v. %v\n", i+1, s)
		}
		fmt.Println("Choose a category: (Or enter * to main menu)")
		input := userInput()
		fmt.Println()
		Cinput, err := strconv.Atoi(input)
		if input == "*" {
			page = 0
			return false
		} else if input == "" {
			fmt.Println("Invalid input!")
			continue
		} else if err != nil {
			fmt.Println("Invalid input!")
			continue
		} else if Cinput <= 0 || Cinput > len(category) {
			fmt.Println("Invalid input!")
			continue
		} else if len(category) == 0 || len(itemsDetail) == 0 || len(items) == 0 {
			fmt.Println("Sorry! Shopping List still not done yet..")
			fmt.Println("Enter to main menu...")
			userInput()
			page = 0
			return false
		}
		tInt = Cinput - 1
		break
	}
	for {
		s, _ := fmt.Printf("Choice from category '%v'\n", category[tInt])
		fmt.Println(strings.Repeat("-", s-1))
		i := 0
		for k, v := range itemsDetail {
			if v.category == tInt {
				i++
				tShopItemName = append(tShopItemName, k)
				fmt.Printf("%v. %v \t$%.2f\n", i, k, v.cost)
			}
		}
		fmt.Println("Pick an item to buy: (Or enter * to main menu)")
		input := userInput()
		Iinput, err := strconv.Atoi(input)
		if input == "*" {
			page = 0
			return false
		} else if input == "" {
			fmt.Println("Invalid input!")
			continue
		} else if err != nil {
			fmt.Println("Invalid input!")
			continue
		} else if Iinput <= 0 || Iinput > i {
			fmt.Println("Invalid input!")
			continue
		}
		tInt = Iinput - 1
		fmt.Println()
		break
	}

	for {
		fmt.Printf("%v) %v\t$%.2f\n", tInt+1, tShopItemName[tInt], itemsDetail[tShopItemName[tInt]].cost)
		fmt.Println("How many quantity to buy? (Or enter * to main menu)")
		input := userInput()
		fmt.Println()
		Iinput, err := strconv.Atoi(input)
		if input == "*" {
			page = 0
			return false
		} else if input == "" {
			fmt.Println("Invalid input!")
			continue
		} else if err != nil {
			fmt.Println("Invalid input!")
			continue
		} else if Iinput <= 0 {
			fmt.Println("Invalid input!")
			continue
		}
		temp := ShoppingList{ItemName: tShopItemName[tInt], Quantity: Iinput}
		shopitemslist = append(shopitemslist, temp)
		break
	}
	// fmt.Println(shopitemslist) // Print test of shopping list record
	return true
}

//Option 2 - Retrieve Shopping List
func retrieveShoppingList() bool {
	var Total float64 = 0.0
	pf := [][]int{{5, 8, 3}}
	var a, b, c, d int
	s, _ := fmt.Println("Retrieve Previous Shopping List")
	fmt.Println(strings.Repeat("-", s-1))
	if len(shopitemslist) == 0 {
		fmt.Println("No record yet!! (enter to main menu)")
		userInput()
		page = 0
		return false
	}
	for _, v := range shopitemslist {
		tpo := []int{}
		temp := itemsDetail[v.ItemName]
		tCName := category[temp.category]
		tQty := strconv.Itoa(v.Quantity)
		tpo = append(tpo, len(v.ItemName), len(tCName), len(tQty))
		pf = append(pf, tpo)
	}
	for i := 0; i < len(pf); i++ {
		if a < pf[i][0] {
			a = pf[i][0]
		}
		if b < pf[i][1] {
			b = pf[i][1]
		}
		if c < pf[i][2] {
			c = pf[i][2]
		}
	}
	fmt.Printf("No.\t")
	fmt.Printf("Items")
	for j := 0; j < (a - 5 + 5); j++ {
		fmt.Printf(" ")
	}
	fmt.Printf("Category")
	for j := 0; j < (b - 8 + 5); j++ {
		fmt.Printf(" ")
	}
	fmt.Printf("Qty")
	for j := 0; j < (c - 3 + 5); j++ {
		fmt.Printf(" ")
	}
	fmt.Printf("Cost\t  ")
	fmt.Printf("Total\t\n")
	for i, v := range shopitemslist {
		temp := itemsDetail[v.ItemName]
		tCName := category[temp.category]
		tQty := v.Quantity
		tPrice := temp.cost
		tTotal := tPrice * float64(tQty)
		fmt.Printf("%v\t", i+1)
		fmt.Printf("%v", v.ItemName)
		for j := 0; j < (a - len(v.ItemName) + 5); j++ {
			fmt.Printf(" ")
		}
		fmt.Printf("%v", tCName)
		for j := 0; j < (b - len(tCName) + 5); j++ {
			fmt.Printf(" ")
		}
		fmt.Printf("%v", tQty)
		for j := 0; j < (c - len(strconv.Itoa(tQty)) + 5); j++ {
			fmt.Printf(" ")
		}
		fmt.Printf("$%.2f\t  ", tPrice)
		fmt.Printf("$%.2f\t\n", tTotal)
		Total = Total + tTotal
		d = i + 1
	}
	fmt.Println()
	fmt.Printf("Total Cost = $%.2f\n", Total)

	for {
		fmt.Println("Delete the shopping list enter the number: (Or enter * to main menu)")
		input := userInput()
		Iinput, err := strconv.Atoi(input)
		if input == "*" {
			page = 0
			return false
		} else if input == "" {
			fmt.Println("Invalid input!")
			continue
		} else if err != nil {
			fmt.Println("Invalid input!")
			continue
		} else if Iinput <= 0 || Iinput > d {
			fmt.Println("Invalid input!")
			continue
		} else {
			shopitemslist = append(shopitemslist[:Iinput-1], shopitemslist[Iinput:]...)
			fmt.Printf("No. %v had deleted.\n", Iinput)
			userInput()
			break
		}
	}
	return true
}

// pull-out function
func userInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func checkData(s string) (bool, int, string) {
	for i, v := range category {
		if s == v {
			return true, i, v
		}
	}
	return false, 0, ""
}

func printFormat(s string, st Item) {

	fmt.Printf("Category: %v - Item: %v Quantity: %v Unit Cost: %v\n", category[st.category], s, st.unit, st.cost)

}

func totalUpValue(index int) (sum float64) {
	for _, v := range itemsDetail {
		if v.category == index {
			sum += v.cost
		}
	}
	return
}

func checkInputModifyValue(i int, j int, s string) (change bool, rs string, d Item) {
	switch i {
	case 1:
		if s == "" {
			change = false
			rs = "No changes to item name made."
			return
		} else if items[j] == s {
			change = false
			rs = "No changes to item name made. It is same as data."
		} else {
			for _, v := range items {
				if s == v {
					change = false
					rs = "The name is already existed! Cannot input again!"
					return
				}
			}
			change = true
			rs = s
			return
		}
	case 2:
		if s == "" {
			change = false
			rs = "No changes to category made."
			return
		} else if category[itemsDetail[items[j]].category] == s {
			change = false
			rs = "No changes to category made. It is same as data."
		} else {
			b, i, _ := checkData(s)
			if b {
				change = true
				d.category = i
				return
			}
			change = false
			rs = "No changes to category made. The category name does not exist."
			return
		}

	case 3:
		if s == "" {
			change = false
			rs = "No changes to quantity made."
			return
		} else {
			v, err := strconv.Atoi(s)
			if err != nil {
				change = false
				rs = "No changes to quantity made. Invalid input!"
				return
			} else if itemsDetail[items[j]].unit == v {
				change = false
				rs = "No changes to quantity made. It is same as data."
			} else {
				change = true
				d.unit = v //need to change return type
				return
			}
		}
	case 4:
		if s == "" {
			change = false
			rs = "No changes to cost made."
			return
		} else {
			v, err := strconv.ParseFloat(s, 10)
			if err != nil {
				change = false
				rs = "No changes to cost made. Invalid input!"
				return
			} else if itemsDetail[items[j]].cost == v {
				change = false
				rs = "No changes to cost made. It is same as data."
			} else {
				change = true
				d.cost = v //need to change return type
				return
			}
		}

	default:
		return
	}
	return
}

func removeCateData(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func funcCategory(input string) bool {

	if input == "" {
		page = 0 // go back to menu
		return false
	} else {
		for _, v := range input {
			if string(v) == " " {
				fmt.Println("Category cannot contain a space(Suggestion - or _)! Enter to continue...")
				userInput()
				return false
			}
		}
		b, i, v := checkData(input)
		if b {
			fmt.Printf("Category: %s already exist at index %v! Enter to continue...\n", v, i)
			userInput()
			return false
		}
	}
	return true
}
