package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	pb "github.com/QuanDN22/BE/gRPC/grpc-gateway-httpbody-message/proto/httpbody_example"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/xuri/excelize/v2"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// // Create a client connection to the gRPC server we just created
	// // This is where the gRPC-gateway proxies the requests

	// approach 1
	// conn, err := grpc.NewClient(
	// 	"0.0.0.0:8080",
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()),
	// )
	// if err != nil {
	// 	log.Fatalln("Failed to dial server:", err)
	// }

	// gwmux := runtime.NewServeMux()
	// // Register Greeter
	// err = helloworldpb.RegisterGreeterHandler(context.Background(), gwmux, conn)
	// if err != nil {
	// 	log.Fatalln("Failed to register gateway:", err)
	// }

	// approach 2
	gwmux := runtime.NewServeMux()

	// Attachment upload from http/s handled manually
	gwmux.HandlePath("POST", "/v1/files", handleBinaryFileUpload)

	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Register Greeter
	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	err := pb.RegisterHttpBodyExampleServiceHandlerFromEndpoint(ctx, gwmux, "0.0.0.0:8083", dialOptions)
	if err != nil {
		log.Fatalln("Failed to register gateway: ", err)
	}

	gwServer := &http.Server{
		Addr:    ":8082",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway is running on http://0.0.0.0:8082")
	log.Fatalln(gwServer.ListenAndServe())
}

func handleBinaryFileUpload(w http.ResponseWriter, r *http.Request, params map[string]string) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse form: %s", err.Error()), http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("attachment")
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get file 'attachment': %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer file.Close()

	fmt.Println(header.Filename, header.Size, header.Header)

	//
	// Now do something with the io.Reader in `f`, i.e. read it into a buffer or stream it to a gRPC client side stream.
	// Also `header` will contain the filename, size etc of the original file.
	//

	// bytes, err := io.ReadAll(f)
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("failed to read file: %s", err.Error()), http.StatusInternalServerError)
	// 	return
	// }
	// read excel document
	f, err := excelize.OpenReader(file)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}()

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		log.Println(err)
		return
	}

	// remove first element in the slice rows
	rows = append(rows[:0], rows[1:]...)

	//create a connection
	conn, err := grpc.NewClient("localhost:8083", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	//create a client
	client := pb.NewHttpBodyExampleServiceClient(conn)

	stream, err := client.Upload(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Upload: %v", err)
	}

	type server struct {
		Server_Names  string
		Server_IPv4   string
		Server_Status string
	}

	// add server in database with three 3 fields: server_name, server_ip, server_status
	for _, row := range rows {
		data, _ := json.Marshal(&server{
			Server_Names:  row[0],
			Server_IPv4:   row[1],
			Server_Status: row[2],
		})

		stream.Send(&pb.Chunk{
			Content: data,
		})
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v\n", err)
	}

	w.Write(res.Data)

	// // buf := bytes.NewBuffer(nil)

	// //create a connection
	// conn, err := grpc.NewClient("localhost:8083", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatalf("did not connect: %v", err)
	// }
	// defer conn.Close()

	// //create a client
	// client := pb.NewHttpBodyExampleServiceClient(conn)

	// stream, err := client.Upload(context.Background())

	// if err != nil {
	// 	log.Fatalf("Error while calling Upload: %v", err)
	// }

	// // buf := make([]byte, 128*1024*1024)
	// // if _, err := io.Copy(buf, bytes); err != nil {
	// // 	return err
	// // }
	// stream.Send(&pb.Chunk{
	// 	Content: bytes,
	// })

	// // writing := true
	// // for writing {
	// // 	// put as many bytes as `chunkSize` into the
	// // 	// buf array.
	// // 	n, err := f.Read(buf)

	// // 	// ... if `eof` --> `writing=false`...
	// // 	if err == io.EOF {
	// // 		writing = false
	// // 	}

	// // 	stream.Send(&pb.Chunk{
	// // 		// because we might've read less than
	// // 		// `chunkSize` we want to only send up to
	// // 		// `n` (amount of bytes read).
	// // 		// note: slicing (`:n`) won't copy the
	// // 		// underlying data, so this as fast as taking
	// // 		// a "pointer" to the underlying storage.
	// // 		Content: buf[:n],
	// // 	})
	// // }

	// res, err := stream.CloseAndRecv()

	// if err != nil {
	// 	log.Fatalf("Error while receiving response from LongGreet: %v\n", err)
	// }

	// // log.Printf("LongGreet: %s\n", res.Data)

	// w.Write(res.Data)
}
