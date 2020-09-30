package main

import (
    "encoding/json"
    "fmt"
)

// interface of base class
type TaskHandler interface {
    InitTask()
}

// interface of derived class
type DerivedTaskHandler interface {
    GetPathOfParam() string
    GetParam() string
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
    *TaskData
    DerivedTaskHandler
}

// member function of base class
// GetPathOfParam is defined in the derived class
func (t Task) InitTask() {
    fmt.Println(t.GetPathOfParam())
    fmt.Println(t.TaskID)
    fmt.Println(t.GetParam())
}

// custom data member of derived class
type TaskAppConfig struct {
    URL string
}

// derived class
type TaskApp struct {
    TaskData
    Config TaskAppConfig
}

// member function of derived class
func (t TaskApp) GetPathOfParam() string {
    return "/edge/app/" + t.TaskID + "/config"
}

// member function of derived class
func (t TaskApp) GetParam() string {
    res, _ := json.Marshal(t.Config)
    return string(res)
}

func main() {
    taskData := TaskData{TaskID: "xxx", Progress: "33", Msg: "wtf", Status: "4"}
    taskAppData := TaskAppConfig{URL: "http://abc.com"}
    taskApp := TaskApp{TaskData: taskData, Config: taskAppData}
    // The three lines above simulate that the constructor of the derived class.
    // As Go doesn't support the real polymorphism, it is Task class, instead of TaskApp class to do the polymorphism-thing.
    task := Task{DerivedTaskHandler: taskApp, TaskData: &taskApp.TaskData}
    task.InitTask()
}
