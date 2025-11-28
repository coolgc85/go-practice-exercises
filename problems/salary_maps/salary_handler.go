package main

import "fmt"

func addEmployee(empMap map[string]float64, employeeName string, salary float64) {
	empMap[employeeName] = salary
	fmt.Printf("Added %s with a salary of %.2f.\n", employeeName, salary)
}

func printAll(empMap map[string]float64) {
	fmt.Println("\n--- Employee Salaries ---")
	for name, salary := range empMap {
		fmt.Printf("Employee: %s, Salary: %.2f\n", name, salary)
	}
	fmt.Println("-----------------------")
}

func main() {
	employeeMap := map[string]float64{
		"Jazzel": 700.0,
		"Audry":  800.0,
		"Kimmy":  9000.0,
	}

	fmt.Println("Initial employee map:")
	printAll(employeeMap)

	keyToLookup := "Victor"
	fmt.Printf("\nChecking for employee: %s...\n", keyToLookup)

	if salary, ok := employeeMap[keyToLookup]; ok {
		fmt.Printf("Salary of %.2f found for %s.\n", salary, keyToLookup)
	} else {
		fmt.Printf("No salary found for %s. Adding them to the map.\n", keyToLookup)
		addEmployee(employeeMap, keyToLookup, 3000)
	}

	keyToRemove := "Jazzel"
	fmt.Printf("\nAttempting to remove employee: %s...\n", keyToRemove)

	if _, ok := employeeMap[keyToRemove]; ok {
		delete(employeeMap, keyToRemove)
		fmt.Printf("Employee %s has been removed.\n", keyToRemove)
	} else {
		fmt.Printf("Employee %s not found, could not remove.\n", keyToRemove)
	}

	fmt.Printf("\nRaising salary for: %s...\n", keyToLookup)
	employeeMap[keyToLookup] += 100.0
	fmt.Printf("New salary is %.2f.\n", employeeMap[keyToLookup])

	fmt.Println("\nFinal employee map:")
	printAll(employeeMap)
}
