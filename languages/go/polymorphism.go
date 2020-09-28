package main

import "fmt"

// interface of derived class
type TaskHandler interface {
        GetPathOfParam() string
}

// class data member
type TaskData struct {
        TaskID   string
        Progress string
        Msg      string
        Status   string
}

// base class containing data member and member function
type Task struct {
        TaskData
        TaskHandler
}

// member function of base class
// GetPathOfParam is defined in the derived class
func (t Task) InitTask() {
        fmt.Println(t.GetPathOfParam())
        fmt.Println(t.TaskID)
}

// custom data member of derived class
type TaskAppConfig struct {
        URL string
}

// derived class
type TaskApp struct {
        *TaskData
        Config TaskAppConfig
}

// member function of derived class
func (t TaskApp) GetPathOfParam() string {
        return "/edge/app/" + t.TaskID + "/config"
}

func main() {
        tData := TaskData{TaskID: "xxx", Progress: "33", Msg: "wtf", Status: "4"}
        t := Task{TaskData: tData, TaskHandler: TaskApp{TaskData: &tData, Config: TaskAppConfig{URL: "xxx"}}}
        t.InitTask()
}
