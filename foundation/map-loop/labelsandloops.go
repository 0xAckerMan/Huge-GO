package main
//lables and loops

func labelandloops(){
    testItemsMap := map [string]int {"ftp":21, "ssh": 22, "smtp": 25 }

    for k := range testItemsMap {
        for _, v := range returnedData {
            if k == v.ID {
                continue outer
            }
        }
        t.Error("Not found")
    }
}
