package chat

import (
    "io/ioutil"
    "os"
    "log"
    "strings"
    "fmt"
    "strconv"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Server struct {
}

var addresses[4] string = [4]string{":9000",":9001",":9002",":9003"}

func (s *Server) SendChunks(ctx context.Context, in *Chunk) (*Signal, error) {
    ioutil.WriteFile(in.Name, in.Data, 0777)
    
    if in.Parts != 0 {
        
    }
    
	return &Signal{Body: "Parte lista"}, nil
}

func (s *Server) TransferenciaLista(ctx context.Context, in *Signal) (*Signal, error) {
    enviar_propuesta(generar_propuesta(int(in.Id)))
    
	return &Signal{Body: "Transferencia terminada"}, nil
}

func (s *Server) EnviarPropuesta(ctx context.Context, in *Signal) (*Signal, error) {
    log.Printf("%s", in.Body)
    
    var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
    
    c := NewChatServiceClient(conn)
    response, err := c.RecibirPropuesta(context.Background(), &Signal{Response: true, Body: in.Body, Id: 1})
    if err != nil {
    log.Fatalf("Error: %s", err)
    }
    log.Printf("%s", response.Body)
    
	return &Signal{Body: "Chunks distribuidos"}, nil
}

func (s *Server) RecibirPropuesta(ctx context.Context, in *Signal) (*Signal, error) {
    log.Printf("Propuesta aprobada")
    
    var current_node = strconv.Itoa(int(in.Id))
    var propuesta = strings.Split(in.Body, ",")
    fmt.Println(propuesta)
    
    for i := 0; i < len(propuesta); i++ {
        if propuesta[i] == current_node {
            continue
        }
        
        
        chunkName := "Frankenstein-Mary_Shelley Parte_"+strconv.Itoa(i+1)
        file, err := os.Open(chunkName)
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
        defer file.Close()
        
        fileInfo, _ := file.Stat()
        var fileSize int64 = fileInfo.Size()
        buffer := make([]byte, fileSize)
        file.Read(buffer)
        
        index, err := strconv.Atoi(propuesta[i])
        if err != nil {
            log.Fatalf("fail: %s", err)
        }
        
        var conn *grpc.ClientConn
        conn, errconn := grpc.Dial(addresses[index], grpc.WithInsecure())
        if errconn != nil {
            log.Fatalf("did not connect: %s", err)
        }
        defer conn.Close()
        
        c := NewChatServiceClient(conn)
        response, err := c.SendChunks(context.Background(), &Chunk{Name: chunkName,Data: buffer})
        if err != nil {
		log.Fatalf("Error: %s", err)
        }
        log.Printf("%s", response.Body)
        
        err = os.Remove(chunkName) 
        if err != nil { 
        log.Fatal(err) 
        } 
    }
    
    
	return &Signal{Body: "Chunks distribuidos"}, nil
}

func generar_propuesta(partes int) string {
    var propuesta string
    var nodo int = 1
    for i := 0; i < partes; i++ {
        if i == partes-1 {
            propuesta += strconv.Itoa(nodo)
            break
        }
        
        propuesta += strconv.Itoa(nodo)+","
        if nodo == 3 {
            nodo = 1
        } else {
            nodo++
        }
    }
    
    return propuesta
}

func enviar_propuesta(propuesta string) {
    var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
    
    c := NewChatServiceClient(conn)
    response, err := c.EnviarPropuesta(context.Background(), &Signal{Id: 1, Body: propuesta})
    if err != nil {
    log.Fatalf("Error: %s", err)
    }
    log.Printf("%s", response.Body)
}
