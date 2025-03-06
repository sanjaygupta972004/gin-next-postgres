
## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/savvy-bit/gin-react-postgres.git
   cd gin-react-postgres
   ```
2. Install `air` to run the app:
   #### Via `go install`
   ```bash
   go install github.com/air-verse/air@latest
   ```
   #### Via `Via install.sh`
   ```bash
   # binary will be $(go env GOPATH)/bin/air
   curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

   # or install it into ./bin/
   curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s

   air -v
   ```
3. Run migration
   ```bash
   go /migrate/migrate.go
   ```
3. Run the server
   ```bash
   air run main.go
   ```

## FAQ
   #### Question: `command not found: air` or `No such file or directory`
   ```bash
   export GOPATH=$HOME/xxxxx
   export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
   export PATH=$PATH:$(go env GOPATH)/bin #Confirm this line in your .profile and make sure to source the .profile if you add it!!!
   ```