package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	myMap := map[string]int{}
	myMap["quarter_of_a_dozen"] = 3
	myMap["half_of_a_dozen"] = 6
	myMap["dozen"] = 12
	myMap["small_gross"] = 120
	myMap["gross"] = 144
	myMap["great_gross"] = 1728
	return myMap
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	// Return false if the given item is not in the bill
	_, exists := units[unit]
	if !exists {
		return false
	}

	// Otherwise add the item to the customer bill, indexed by the item name, then return true.
	// If the item is already present in the bill, increase its quantity by the amount that belongs to the provided unit.

	bill[item] += units[unit]
	return true

	// // Return false if the given unit is not in the units map
	// myMeasure, exists := units[unit]
	// if !exists {
	// 	return false
	// }

	// Return false if the new quantity would be less than 0.

	// If the new quantity is 0, completely remove the item from the bill then return true.
	// Otherwise, reduce the quantity of the item and return true.

}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	var exists bool

	// Return false if the given item is not in the bill
	_, exists = bill[item]
	if !exists {
		return false
	}

	// Return false if the given unit is not in the units map.
	_, exists = units[unit]
	if !exists {
		return false
	}

	// Return false if the new quantity would be less than 0.
	if bill[item] < units[unit] {
		return false
	}

	// If the new quantity is 0, completely remove the item from the bill then return true.
	if bill[item] == units[unit] {
		delete(bill, item)
		return true
	}

	// Otherwise, reduce the quantity of the item and return true.
	bill[item] -= units[unit]
	return true

}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	//Return 0 and false if the item is not in the bill.
	//Otherwise, return the quantity of the item in the bill and true.
	v, exists := bill[item]
	return v, exists
}
