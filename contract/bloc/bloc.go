package main


// 2. 외부모듈 포함
import (
   "bytes"
   "encoding/json"
   "fmt"
   "strconv"
   "time"

   "github.com/hyperledger/fabric/core/chaincode/shim"
   "github.com/hyperledger/fabric/protos/peer"
)

 

// 3. 체인코드 클래스 정의

type SimpleAsset struct {
}

 

// 4. 구조체 정의
type Student struct {

   Id               string     `json:"id"`
   Attandance       Attandance `json:"attandance"`
   Current_state    string     `json:"cureent-state"` // { FREE, ENTERED, SITTED, MAT-DOWNLOADED, NOTE-UPDATED, QUALIFIED }
   Total_count      int        `json:"total_count"`   // 출석율
   Attandance_count int        `json:"attandance_count"`
   Attandance_rate  float64    `json:"attandance_rate"`
}

type Attandance struct {
   Class_name              string   `json:"class_name"`
   Enter_time              string   `json:"enter_time"`
   Sit_no                  string   `json:"sit_no"`
   Class_material_download bool     `json:"class_material_download"`
   Query_answer_list       []string `json:"query_answer_list"`
   Lecture_note            string   `json:"lecture_note"`
   Exit_time               string   `json:"exit_time"`
}

// 5. Init 함수 정의

func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {

   return shim.Success([]byte("Init process was done."))

}

 

// 6. Invoke 함수 정의

func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

   fn, args := stub.GetFunctionAndParameters()
   var result string
   var err error
   if fn == "addStudent" {
      result, err = addStudent(stub, args)
   } else if fn == "attand" {
      result, err = attand(stub, args)
   } else if fn == "sit" {
      result, err = sit(stub, args)
   } else if fn == "download_material" {
      result, err = download_material(stub, args)
   } else if fn == "query_answer" {
      result, err = query_answer(stub, args)
   } else if fn == "save_note" {
      result, err = save_note(stub, args)
   } else if fn == "exit" { // 출석률 갱신
      result, err = exit(stub, args)
   } else if fn == "queryStudent" {
      result, err = queryStudent(stub, args)
   } else if fn == "queryStudentHistory" {
      result, err = queryStudentHistory(stub, args)
   } else {
      return shim.Error(err.Error())
   }
   if err != nil {
      return shim.Error(err.Error())
   }

   return shim.Success([]byte(result))
}

// 7. AddStudent 함수 정의

func addStudent(stub shim.ChaincodeStubInterface, args []string) (string, error) {

   if len(args) != 2 {
      return "", fmt.Errorf("Incorrect arguments. Expecting a key and a value")
   }

   fmt.Printf("addStudent: start")
   fmt.Println("set PutState: " + args[0])
   total_count, _ := strconv.Atoi(args[1])

   var data = Student{Id: args[0], Current_state: "FREE", Total_count: total_count, Attandance_count: 0 /*, Attandance_rate: 0.0*/}
   dataAsBytes, _ := json.Marshal(data)
   err := stub.PutState(args[0], dataAsBytes)
   if err != nil {

      return "", fmt.Errorf(fmt.Sprintf("Failed to create asset: %s", args[0]))

   }

   return args[0], nil

}

 

// 8. Attand 함수 정의

func attand(stub shim.ChaincodeStubInterface, args []string) (string, error) {

   if len(args) != 2 {

      return "", fmt.Errorf("Incorrect arguments. Expecting a key")

   }

   fmt.Println("get GetState: " + args[0])

   value, err := stub.GetState(args[0])

   if err != nil { // GetState 함수가 오류 난 경우

      return "", fmt.Errorf(fmt.Sprintf("Failed to create asset: %s", args[0]), err)

   }

   if value == nil { // key에 대한 값을 찾을 수 없는 경우

      return "", fmt.Errorf("Asset not found: %s", args[0])

   }

   data := Student{}

   json.Unmarshal(value, &data)

   data.Current_state = "ENTERED"

 

   data.Attandance.Class_name = args[1]

   data.Attandance.Enter_time = time.Now().String()

   dataAsBytes, _ := json.Marshal(data)

   err = stub.PutState(args[0], dataAsBytes)

   if err != nil {

      return "", fmt.Errorf("Failed to set asset: %s", args[0])

   }

   return string(dataAsBytes), nil

}

 

// 8. Sit 함수 정의

func sit(stub shim.ChaincodeStubInterface, args []string) (string, error) {

   if len(args) != 3 {

      return "", fmt.Errorf("Incorrect arguments. Expecting a key")

   }

   fmt.Println("get GetState: " + args[0])

   value, err := stub.GetState(args[0])

   if err != nil { // GetState 함수가 오류 난 경우

      return "", fmt.Errorf(fmt.Sprintf("Failed to create asset: %s", args[0]), err)

   }

   if value == nil { // key에 대한 값을 찾을 수 없는 경우

      return "", fmt.Errorf("Asset not found: %s", args[0])

   }

   data := Student{}

   json.Unmarshal(value, &data)

   if data.Attandance.Class_name == args[1] {

      data.Current_state = "SITTED"

      data.Attandance.Sit_no = args[2]

   }

   dataAsBytes, _ := json.Marshal(data)

   err = stub.PutState(args[0], dataAsBytes)

   if err != nil {

      return "", fmt.Errorf("Failed to set asset: %s", args[0])

   }

   return string(dataAsBytes), nil

}

 

// 8. Download_material 함수 정의

func download_material(stub shim.ChaincodeStubInterface, args []string) (string, error) {

   if len(args) != 2 {

      return "", fmt.Errorf("Incorrect arguments. Expecting a key")

   }

   fmt.Println("get GetState: " + args[0])

   value, err := stub.GetState(args[0])

   if err != nil { // GetState 함수가 오류 난 경우

      return "", fmt.Errorf(fmt.Sprintf("Failed to create asset: %s", args[0]), err)

   }

   if value == nil { // key에 대한 값을 찾을 수 없는 경우

      return "", fmt.Errorf("Asset not found: %s", args[0])

   }

   data := Student{}

   json.Unmarshal(value, &data)

   if data.Attandance.Class_name == args[1] {

      data.Current_state = "MT_DOWNLOADED"

      data.Attandance.Class_material_download = true

   }

   dataAsBytes, _ := json.Marshal(data)

   err = stub.PutState(args[0], dataAsBytes)

   if err != nil {

      return "", fmt.Errorf("Failed to set asset: %s", args[0])

   }

   return string(dataAsBytes), nil

}

 

// 8. query_answer 함수 정의

func query_answer(stub shim.ChaincodeStubInterface, args []string) (string, error) {

   if len(args) != 4 {

      return "", fmt.Errorf("Incorrect arguments. Expecting a key")

   }

   fmt.Println("get GetState: " + args[0])

   value, err := stub.GetState(args[0])

   if err != nil { // GetState 함수가 오류 난 경우

      return "", fmt.Errorf(fmt.Sprintf("Failed to create asset: %s", args[0]), err)

   }

   if value == nil { // key에 대한 값을 찾을 수 없는 경우

      return "", fmt.Errorf("Asset not found: %s", args[0])

   }

   data := Student{}

   json.Unmarshal(value, &data)

   if data.Attandance.Class_name == args[1] {

      data.Attandance.Query_answer_list = append(data.Attandance.Query_answer_list, "{query:\""+args[2]+"\"answer:\""+args[3]+"\"}")

   }

   dataAsBytes, _ := json.Marshal(data)

   err = stub.PutState(args[0], dataAsBytes)

   if err != nil {

      return "", fmt.Errorf("Failed to set asset: %s", args[0])

   }

   return string(dataAsBytes), nil

}

 

// 8. save_note 함수 정의

func save_note(stub shim.ChaincodeStubInterface, args []string) (string, error) {

   if len(args) != 3 {

      return "", fmt.Errorf("Incorrect arguments. Expecting a key")

   }

   fmt.Println("get GetState: " + args[0])

   value, err := stub.GetState(args[0])

   if err != nil { // GetState 함수가 오류 난 경우

      return "", fmt.Errorf(fmt.Sprintf("Failed to create asset: %s", args[0]), err)

   }

   if value == nil { // key에 대한 값을 찾을 수 없는 경우

      return "", fmt.Errorf("Asset not found: %s", args[0])

   }

   data := Student{}

   json.Unmarshal(value, &data)

   if data.Attandance.Class_name == args[1] {

      data.Current_state = "NOTE_UPDATE"

      data.Attandance.Lecture_note = args[2]

   }

   dataAsBytes, _ := json.Marshal(data)

   err = stub.PutState(args[0], dataAsBytes)

   if err != nil {

      return "", fmt.Errorf("Failed to set asset: %s", args[0])

   }

   return string(dataAsBytes), nil

}

 

// 8. exit 함수 정의

func exit(stub shim.ChaincodeStubInterface, args []string) (string, error) {

   if len(args) != 2 {

      return "", fmt.Errorf("Incorrect arguments. Expecting a key")

   }

   fmt.Println("get GetState: " + args[0])

   value, err := stub.GetState(args[0])

   if err != nil { // GetState 함수가 오류 난 경우

      return "", fmt.Errorf(fmt.Sprintf("Failed to create asset: %s", args[0]), err)

   }

   if value == nil { // key에 대한 값을 찾을 수 없는 경우

      return "", fmt.Errorf("Asset not found: %s", args[0])

   }

   data := Student{}

   json.Unmarshal(value, &data)

 

   fmt.Print("exit:" + data.Attandance.Class_name + ":" + args[1])

 

   if data.Attandance.Class_name == args[1] {

 

      data.Current_state = "QUALIFIED"

 

      data.Attandance_count++

      data.Attandance_rate = float64(data.Attandance_count) / float64(data.Total_count)

 

      data.Attandance.Exit_time = time.Now().String()

   }

 

   dataAsBytes, _ := json.Marshal(data)

   err = stub.PutState(args[0], dataAsBytes)

   if err != nil {

      return "", fmt.Errorf("Failed to set asset: %s", args[0])

   }

   return string(dataAsBytes), nil

}

func queryStudent(stub shim.ChaincodeStubInterface, args []string) (string, error) {

   if len(args) != 1 {

      return "", fmt.Errorf("Incorrect arguments. Expecting a key")

   }

   value, err := stub.GetState(args[0])

   if err != nil {

      return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)

   }

   if value == nil {

      return "", fmt.Errorf("Asset not found: %s", args[0])

   }

   return string(value), nil

}

func queryStudentHistory(stub shim.ChaincodeStubInterface, args []string) (string, error) {

   // 전달 인자 체크 args ? key

   if len(args) < 1 {

      return "", fmt.Errorf("Incorrect number of arguments. Expecting 1")

   }

   keyName := args[0]

   // getHistoryForKey 함수가 시작되는 로그를 남기기

   fmt.Printf("- start getHistoryForMarble: %s\n", keyName)

   resultsIterator, err := stub.GetHistoryForKey(keyName)

   if err != nil {

      return "", err

   }

   defer resultsIterator.Close()

   // buffer is a JSON array containing historic values for key

   var buffer bytes.Buffer

   buffer.WriteString("[")

   bArrayMemberAlreadyWritten := false

   for resultsIterator.HasNext() {

      // resultsIterator에서 next 아이템 꺼내서 response 할당

      response, err := resultsIterator.Next()

      if err != nil {

         return "", err

      }

      // Add a comma before array members, suppress it for the first array member

      if bArrayMemberAlreadyWritten == true {

         buffer.WriteString(",")

      }

      buffer.WriteString("{\"TxId\":")

      buffer.WriteString("\"")

      buffer.WriteString(response.TxId)

      buffer.WriteString("\"")

      buffer.WriteString(", \"Value\":")

      // if it was a delete operation on given key, then we need to set the

      //corresponding value null. Else, we will write the response.Value

      //as-is (as the Value itself a JSON marble)

      if response.IsDelete {

         buffer.WriteString("null")

      } else {

         buffer.WriteString(string(response.Value))

      }

      buffer.WriteString(", \"Timestamp\":")

      buffer.WriteString("\"")

      buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())

      buffer.WriteString("\"")

      buffer.WriteString(", \"IsDelete\":")

      buffer.WriteString("\"")

      buffer.WriteString(strconv.FormatBool(response.IsDelete))

      buffer.WriteString("\"")

      buffer.WriteString("}")

      bArrayMemberAlreadyWritten = true

   }

   buffer.WriteString("]")

   fmt.Printf("- getHistoryForKey returning:\n%s\n", buffer.String())

   return (string)(buffer.Bytes()), nil

}

 

// 9. main 함수 정의

func main() {

   if err := shim.Start(new(SimpleAsset)); err != nil {

      fmt.Printf("Error starting SimpleAsset chaincode: %s", err)

   }

}