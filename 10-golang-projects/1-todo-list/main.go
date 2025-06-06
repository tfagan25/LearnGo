package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

/*
To Do List App
--------------
- multiple named to do lists
- create a new list
- show items in single to do list
- add item to list
- remove item from list
- write to file
*/

type list struct {
	name  string
	items []string
}

func showAllLists(allLists []list) {
	fmt.Printf("All Lists\n---------\n")
	for idx, list := range allLists {
		fmt.Printf("%d: %s\n", idx+1, list.name)
	}
	fmt.Println()
}

func createNewList(allLists *[]list, listName string) {
	*allLists = append(*allLists, list{name: listName})
}

func removeList(allLists *[]list, listName string) {
	for i := range *allLists {
		if (*allLists)[i].name == listName {
			(*allLists) = slices.Delete((*allLists), i, i + 1)
			return
		}
	}
}

func (lst list) showList() {
	fmt.Printf("%s\n%s\n", lst.name, strings.Repeat("-", len(lst.name)))
	if len(lst.items) > 0 {
		for idx, item := range lst.items {
			fmt.Printf("%d: %s\n", idx + 1, item)
		}
	} else {
		fmt.Printf("List is empty.\n")
	}
	fmt.Println()
}

func addToList(allLists *[]list, listName string, listItem string) {
	for i := range *allLists {
		if (*allLists)[i].name == listName {
			(*allLists)[i].items = append((*allLists)[i].items, listItem)
			return
		}
	}
}

func removeFromList(allLists *[]list, listName string, listItem string) {
	for i := range *allLists {
		if (*allLists)[i].name == listName {
			removalIdx := 0
			// identify index
			for j := range (*allLists)[i].items {
				if (*allLists)[i].items[j] == listItem {
					removalIdx = j
					break
				}
			}

			(*allLists)[i].items = slices.Delete((*allLists)[i].items, removalIdx, removalIdx + 1)
			return
		}
	}
}

func outputHeading() {
	fmt.Println("---------------------")
	fmt.Println("    TO DO LIST")
	fmt.Println("---------------------")
	fmt.Println("1. Show All Lists")
	fmt.Println("2. Create New List")
	fmt.Println("3. Delete List")
	fmt.Println("4. Show List")
	fmt.Println("5. Add Item to List")
	fmt.Println("6. Remove Item from List")
	fmt.Println()
	fmt.Print("> ")
}

func promptUser() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(input)
}

func main() {
	lists := make([]list, 0)
	
	for {
		outputHeading()
		input := promptUser()
		fmt.Println()

		switch input {
		case "1":
			showAllLists(lists)
		case "2":
			fmt.Print("New List Name: ")
			input := promptUser()
			createNewList(&lists, input)
			fmt.Println()
		case "3":
			fmt.Print("Delete List Name: ")
			input := promptUser()
			removeList(&lists, input)
			fmt.Println()
		case "4":
			fmt.Print("Show List Name: ")
			input := promptUser()
			fmt.Println()

			for _, lst := range lists {
				if lst.name == input {
					lst.showList()
				}
			}
		case "5":
			fmt.Print("List to Add To: ")
			list := promptUser()

			fmt.Print("New List Item: ")
			newItem := promptUser()

			addToList(&lists, list, newItem)
			fmt.Println()
		case "6":
			fmt.Print("List to Remove From: ")
			list := promptUser()

			fmt.Print("Remove List Item: ")
			removeItem := promptUser()

			removeFromList(&lists, list, removeItem)
			fmt.Println()
		default:
			fmt.Println("Please select a list option.")
		}
	}
}
