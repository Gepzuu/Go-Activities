package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func main() {
	// Create a new Fyne application
	myApp := app.New()
	myWindow := myApp.NewWindow("GWA Calculator")

	// Create input fields for grades and units
	gradeEntry := widget.NewEntry()
	gradeEntry.SetPlaceHolder("Enter grades (comma separated)")

	unitEntry := widget.NewEntry()
	unitEntry.SetPlaceHolder("Enter units (comma separated)")

	// Create a label to display the result
	resultLabel := widget.NewLabel("")

	// Create a button to calculate the GWA
	calculateButton := widget.NewButton("Calculate GWA", func() {
		grades := gradeEntry.Text
		units := unitEntry.Text

		// Split the input strings and convert to float64
		gradeList := splitAndConvert(grades)
		unitList := splitAndConvert(units)

		// Calculate the GWA
		gwa := calculateGWA(gradeList, unitList)

		// Display the result
		resultLabel.SetText("GWA: " + strconv.FormatFloat(gwa, 'f', 2, 64))
	})

	// Create a container and add the UI elements
	content := container.NewVBox(
		widget.NewLabel("GWA Calculator"),
		gradeEntry,
		unitEntry,
		calculateButton,
		resultLabel,
	)

	// Set the content of the window
	myWindow.SetContent(content)

	// Show the window and run the application
	myWindow.ShowAndRun()
}

// Helper function to split a comma-separated string and convert to float64
func splitAndConvert(input string) []float64 {
	var result []float64
	for _, s := range strings.Split(input, ",") {
		if val, err := strconv.ParseFloat(strings.TrimSpace(s), 64); err == nil {
			result = append(result, val)
		}
	}
	return result
}

// Function to calculate the GWA
func calculateGWA(grades, units []float64) float64 {
	var totalGradeUnits, totalUnits float64
	for i := range grades {
		totalGradeUnits += grades[i] * units[i]
		totalUnits += units[i]
	}
	return totalGradeUnits / totalUnits
}