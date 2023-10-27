# My Project ASCII Art WEB

## Description
This project is an ASCII Art application that allows you to generate ASCII art using different banner styles. Users can enter text and choose a banner style, then the app generates corresponding ASCII art.

## Authors
- Maodo LO (https://learn.zone01dakar.sn/git/maolo)
- Serigne Fallou DIA (https://learn.zone01dakar.sn/git/sedia)
- Malang COULIBALY (https://learn.zone01dakar.sn/git/malcoulibaly)

## Use

The application uses the following algorithm to generate ASCII art:

The user enters text in the text box.
The application reads the banner file corresponding to the style chosen by the user.
Using the banner dictionary and formatted text, the application generates ASCII art by replacing each character in the text with its representation in the banner.
The generated ASCII art is displayed to the user.
For more details on the source code, please see the main.go file and the asciifs package.

### How to Run
1. Make sure you have Go installed on your system.
2. Clone the project repository: git clone https://learn.zone01dakar.sn/git/sedia/ascii-art-web
3. Navigate to the project directory: cd ascii-art-web
4. Run the application using the following command: go run main.go
5. Open your web browser and navigate to http://localhost:8080 to access the application user interface.

## Example of use
Enter the text you want to convert to ASCII art in the text area.
Choose the banner style from the options (Standard, Shadow, ThinkerToy).
Click the “Generate ASCII Art” button.
The generated ASCII art will be displayed on the screen.