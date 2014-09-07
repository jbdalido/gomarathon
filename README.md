Go client for marathon
======================

[![GoDoc](https://godoc.org/github.com/jbdalido/gomarathon?status.svg)](https://godoc.org/github.com/jbdalido/gomarathon)

Versions:

 - **2014/08/30 Tag V0.1.0**: 
   Fixed version compatible with marathon 0.6 

Dev :

 - **Branch marathon-0.7**: 
   Dev version for marathon 0.7

**_* This is a BETA_**

This is a small go client to use with **[Marathon api v2](https://mesosphere.github.io/marathon/docs/rest-api.html)**. All endpoints are handled (EventSubscriptions have not yet been tested).

Feel free to hack.

### TODO

* Tests like everywhere
* Params building might be much more clean

## Usage example
================
    
    package main

    import (
        "encoding/json"
        "github.com/jbdalido/gomarathon"
        "log"
    )

    func main() {

        c, err := gomarathon.NewClient("http://example.io:8080", nil)
        if err != nil {
            log.Fatal(err)
        }

        // Update app
        a := &gomarathon.Application{
            Id:  "test_app",
            Mem: 515,
            Container: &gomarathon.Container{
                Image:   "docker://jbaptiste/envspitter",
                Options: []string{"-p", "1314:8080"},
            },
        }
        r, err := c.CreateApp(a)
        if err != nil {
            log.Fatal(err)
        }
        v, _ := json.Marshal(r)
        log.Printf("%s", v)

        // List all apps
        r, err = c.ListApps()
        if err != nil {
            log.Fatal(err)
        }
        v, _ = json.Marshal(r)
        log.Printf("%s", v)

        // List all apps
        r, err = c.ListApps()
        if err != nil {
            log.Fatal(err)
        }
        v, _ = json.Marshal(r)
        log.Printf("%s", v)

        // List one app
        r, err = c.GetApp("test_app")
        if err != nil {
            log.Fatal(err)
        }
        v, _ = json.Marshal(r)
        log.Printf("%s", v)

        // List Versions
        r, err = c.ListAppVersions("test_app")
        if err != nil {
            log.Fatal(err)
        }
        v, _ = json.Marshal(r)
        log.Printf("%s", v)

        // Update app
        a = &gomarathon.Application{
            Mem:       515,
            Instances: 2,
        }
        r, err = c.UpdateApp("test_app", a)
        if err != nil {
            log.Fatal(err)
        }
        v, _ = json.Marshal(r)
        log.Printf("%s", v)

        // Get all tasks
        r, err = c.ListTasks()
        if err != nil {
            log.Fatal(err)
        }
        v, _ = json.Marshal(r)
        log.Printf("%s", v)

        // Delete app
        r, err = c.DeleteApp("test_app")
        if err != nil {
            log.Fatal(err)
        }
        v, _ = json.Marshal(r)
        log.Printf("%s", v)

    }

## Authors
==========
- Jean-Baptiste Dalido <jbdalido@gmail.com>
- Anton Lindström <https://github.com/antonlindstrom>
