//I dont remember the comamnds and I need a dox ima forget the pipeline too quick... Note to self have to rewrite data pipeline and figure and how to write a wireless modem
//*********************************************************Header******************************************************************************************

export GODEBUG=netdns=cgo+1

//docs
package main

import("fmt"
       "os"
       "net/ip"
       "net")

func main(){

    ln, err := net.Listen("ip",/*"addr()"*/)
    if err != nil{
        //err.handlr()
    }
    for{
        conn,err := ln.Accept()
        if err != nil{
            //err.Handlr()
        }
        go handleConnection(conn)//I think you can rewrite stuff line conn and handleConnection
        
    }
    
}

//Finna write config files to import what I need. This server messy alrdy
