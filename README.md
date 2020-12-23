# multi
Golang separate on multiple files

Steps:
===
1. Clone or initialize empty git project
   - Example: git clone https://github.com/egevorkyan/multi.git
2. Enter to cloned project folder and prepare project structure
   - Example: project multi will have 2 packages: employee and sources
3. Enter following commang to initialize
   - Example: go mod init github.com/egevorkyan/multi
   - Example Output: go: creating new go.mod: module github.com/egevorkyan/multi
4. Finished tructure of my project created, time to write code
5. main function must be placed in root folder of your project
   - Example: empsalary.go, locates on same level as go.mod
6. Simplefile employees.txt created which contains below information (Full Name, basic salary, tax)
   Edi Tester,100000,33
   Ben Parcker,1500000,40
   Bill Nelson,45000,3
7. Result - it can be any name of file and location, application will ask to input source file
   Enter Employees source file: /Users/eduardgevorkyan/Documents/GOLEARN/employees.txt
   Employee Id:     0
   Employee Name:   Eduard Gevorkyan
   Employee Salary:         67000
   Employee Id:     1
   Employee Name:   Ben Parcker
   Employee Salary:         900000
   Employee Id:     2
   Employee Name:   Bill Nelson
   Employee Salary:         43650
