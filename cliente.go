 package main

 import (
         "fmt"
         "log"
         "io/ioutil"
         "math"
         "os"
         "strconv"
         "bufio"
         "context"
         
         "github.com/nchcl/sd/chat"
         "google.golang.org/grpc"
 )

func send_chunk() {
    var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

    fileToBeChunked := "Frankenstein-Mary_Shelley.pdf"
    file, err := os.Open(fileToBeChunked)
    if err != nil {
            fmt.Println(err)
            os.Exit(1)
    }
    defer file.Close()

    fileInfo, _ := file.Stat()
    var fileSize int64 = fileInfo.Size()
    const fileChunk = 256000 // 1 MB, change this to your requirement
    totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))
    
    for i := uint64(0); i < totalPartsNum; i++ {
        partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
        partBuffer := make([]byte, partSize)
        file.Read(partBuffer)
        
        fileName := "Frankenstein-Mary_Shelley Parte_" + strconv.FormatUint(i+1, 10)
        
        c := chat.NewChatServiceClient(conn)
        response, err := c.SendChunks(context.Background(), &chat.Chunk{Name: fileName,Parts: int32(totalPartsNum),Data: partBuffer})
        if err != nil {
		log.Fatalf("Error: %s", err)
        }
        log.Printf("%s", response.Body)
    }
    
    c := chat.NewChatServiceClient(conn)
    response, err := c.TransferenciaLista(context.Background(), &chat.Signal{Id: int32(totalPartsNum)})
    if err != nil {
    log.Fatalf("Error: %s", err)
    }
    log.Printf("%s", response.Body)
    
	
}
 
func uploader(nombre_libro string){
        fileToBeChunked := nombre_libro

        file, err := os.Open(fileToBeChunked)

        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }

        defer file.Close()

        fileInfo, _ := file.Stat()

        var fileSize int64 = fileInfo.Size()
        var ip string
        const fileChunk = 256000 // 1 MB, change this to your requirement

        // calculate total number of parts the file will be chunked into

        totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

        fmt.Printf("%s | %d \n", fileToBeChunked, totalPartsNum)

        for i := uint64(0); i < totalPartsNum; i++ {
                ip=strconv.FormatUint(i+1, 10) //AQUI DEBE IR LA IP DONDE SE GUARDARA EL CHUNK

                partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
                partBuffer := make([]byte, partSize)

                file.Read(partBuffer)

                // write to disk
                fileName := "Parte_" + strconv.FormatUint(i+1, 10)
                _, err := os.Create(fileName)

                if err != nil {
                        fmt.Println(err)
                        os.Exit(1)
                }

                // write/save buffer to disk
                ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)

                fmt.Println(fileName, " | ", ip)
        }
}


func downloader(nombre_libro string){

        fileToBeChunked := nombre_libro // change here!

         file, err := os.Open(fileToBeChunked)

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         defer file.Close()

         fileInfo, _ := file.Stat()

         var fileSize int64 = fileInfo.Size()

         const fileChunk = 1 * (1 << 20) // 1 MB, change this to your requirement

         // calculate total number of parts the file will be chunked into

         totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

         fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)

         for i := uint64(0); i < totalPartsNum; i++ {

                 partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
                 partBuffer := make([]byte, partSize)

                 file.Read(partBuffer)

                 // write to disk
                 fileName := "bigfile_" + strconv.FormatUint(i, 10)
                 _, err := os.Create(fileName)

                 if err != nil {
                         fmt.Println(err)
                         os.Exit(1)
                 }

                 // write/save buffer to disk
                 ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)

                 fmt.Println("Split to : ", fileName)
         }

         // just for fun, let's recombine back the chunked files in a new file

         newFileName := nombre_libro
         _, err = os.Create(newFileName)

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         //set the newFileName file to APPEND MODE!!
         // open files r and w

         file, err = os.OpenFile(newFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         // IMPORTANT! do not defer a file.Close when opening a file for APPEND mode!
         // defer file.Close()

         // just information on which part of the new file we are appending
         var writePosition int64 = 0

         for j := uint64(0); j < totalPartsNum; j++ {

                 //read a chunk
                 currentChunkFileName := "bigfile_" + strconv.FormatUint(j, 10)

                 newFileChunk, err := os.Open(currentChunkFileName)

                 if err != nil {
                         fmt.Println(err)
                         os.Exit(1)
                 }

                 defer newFileChunk.Close()

                 chunkInfo, err := newFileChunk.Stat()

                 if err != nil {
                         fmt.Println(err)
                         os.Exit(1)
                 }

                 // calculate the bytes size of each chunk
                 // we are not going to rely on previous data and constant

                 var chunkSize int64 = chunkInfo.Size()
                 chunkBufferBytes := make([]byte, chunkSize)

                 fmt.Println("Appending at position : [", writePosition, "] bytes")
                 writePosition = writePosition + chunkSize

                 // read into chunkBufferBytes
                 reader := bufio.NewReader(newFileChunk)
                 _, err = reader.Read(chunkBufferBytes)

                 if err != nil {
                         fmt.Println(err)
                         os.Exit(1)
                 }

                 // DON't USE ioutil.WriteFile -- it will overwrite the previous bytes!
                 // write/save buffer to disk
                 //ioutil.WriteFile(newFileName, chunkBufferBytes, os.ModeAppend)

                 n, err := file.Write(chunkBufferBytes)

                 if err != nil {
                         fmt.Println(err)
                         os.Exit(1)
                 }

                 file.Sync() //flush to disk

                 // free up the buffer for next cycle
                 // should not be a problem if the chunk size is small, but
                 // can be resource hogging if the chunk size is huge.
                 // also a good practice to clean up your own plate after eating

                 chunkBufferBytes = nil // reset or empty our buffer

                 fmt.Println("Written ", n, " bytes")

                 fmt.Println("Recombining part [", j, "] into : ", newFileName)
         }

         // now, we close the newFileName
         file.Close()
        /*fmt.Println(nombre_libro) 
        newFileName := nombre_libro+".zip"
        _, err = os.Create(newFileName)

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         //set the newFileName file to APPEND MODE!!
         // open files r and w

         file, err = os.OpenFile(newFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         // IMPORTANT! do not defer a file.Close when opening a file for APPEND mode!
         // defer file.Close()

         // just information on which part of the new file we are appending
         var writePosition int64 = 0
         var totalPartsNum=5 //Aqui deberia indicarse la cantidad de partes que se dividio el archivo

         for j := 0; j < totalPartsNum; j++ {

                 //read a chunk
                 currentChunkFileName := "Parte_" + strconv.Itoa(j)

                 newFileChunk, err := os.Open(currentChunkFileName)

                 if err != nil {
                         fmt.Println(err)
                         os.Exit(1)
                 }

                 defer newFileChunk.Close()

                 chunkInfo, err := newFileChunk.Stat()

                 if err != nil {
                         fmt.Println(err)
                         os.Exit(1)
                 }

                 // calculate the bytes size of each chunk
                 // we are not going to rely on previous data and constant

                 var chunkSize int64 = chunkInfo.Size()
                 chunkBufferBytes := make([]byte, chunkSize)

                 fmt.Println("Appending at position : [", writePosition, "] bytes")
                 writePosition = writePosition + chunkSize

                 // read into chunkBufferBytes
                 reader := bufio.NewReader(newFileChunk)
                 _, err = reader.Read(chunkBufferBytes)

                 if err != nil {
                         fmt.Println(err)
                         os.Exit(1)
                 }

                 // DON't USE ioutil.WriteFile -- it will overwrite the previous bytes!
                 // write/save buffer to disk
                 //ioutil.WriteFile(newFileName, chunkBufferBytes, os.ModeAppend)

                 n, err := file.Write(chunkBufferBytes)

                 if err != nil {
                         fmt.Println(err)
                         os.Exit(1)
                 }

                 file.Sync() //flush to disk

                 // free up the buffer for next cycle
                 // should not be a problem if the chunk size is small, but
                 // can be resource hogging if the chunk size is huge.
                 // also a good practice to clean up your own plate after eating

                 chunkBufferBytes = nil // reset or empty our buffer

                 fmt.Println("Written ", n, " bytes")

                 fmt.Println("Recombining part [", j, "] into : ", newFileName)
         }

         // now, we close the newFileName
         file.Close()*/

}

func main() {

        var tipo int
        var nombre_libro string
    
        fmt.Println("1. Uploader")
        fmt.Println("2. Downloader")
        fmt.Println("3. Test")
        fmt.Scan(&tipo)

        switch tipo {
                case 1:
                        fmt.Println("¿Que libro desea subir?")
                        fmt.Scan(&nombre_libro)                 
                        uploader(nombre_libro)
                case 2:
                        fmt.Println("¿Que libro desea descargar?")
                        fmt.Scan(&nombre_libro)   
                        downloader(nombre_libro)
                case 3:
                        send_chunk()
        }

        
 }
