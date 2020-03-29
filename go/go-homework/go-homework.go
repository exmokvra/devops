package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.etcd.io/etcd/clientv3"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/calc/sum/{valA}/{valB}", sum)
	router.HandleFunc("/calc/subtract/{valA}/{valB}", subtract)
	router.HandleFunc("/calc/multiply/{valA}/{valB}", multiply)
	router.HandleFunc("/calc/divide/{valA}/{valB}", divide)
	router.HandleFunc("/calc/history", history)
	router.HandleFunc("/calc/deleteAllUserData", deleteUserData)

	http.ListenAndServe(":8080", router)
}

func deleteUserData(w http.ResponseWriter, r *http.Request) {
	cli, ctx, cancel, err := etcdConnect()
	if err != nil {
		fmt.Println("Error starting connection")
	}

	opts := []clientv3.OpOption{
		clientv3.WithPrefix(),
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend),
	}

	result, getErr := listHistoryOperations(cli, ctx, opts)
	if getErr != nil {
		fmt.Println("Error listing operations")
	}
	for _, item := range result.Kvs {
		delResp, delErr := cli.Delete(ctx, string(item.Key))
		if delErr != nil {
			fmt.Println("Error deleting keys")
		} else {
			fmt.Printf("%d keys deleted succesfully", delResp.Deleted)
			fmt.Println()
		}
	}

	closeEtcd(cli, cancel)
}

func history(w http.ResponseWriter, r *http.Request) {
	cli, ctx, cancel, err := etcdConnect()
	if err != nil {
		fmt.Println("Error starting connection")
	}

	opts := []clientv3.OpOption{
		clientv3.WithPrefix(),
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend),
	}

	result, getErr := listHistoryOperations(cli, ctx, opts)
	if getErr != nil {
		fmt.Println("Error listing operations")
	}
	for _, item := range result.Kvs {
		fmt.Println(string(item.Key), string(item.Value))
	}

	closeEtcd(cli, cancel)
}

func listHistoryOperations(cli *clientv3.Client, ctx context.Context, opts []clientv3.OpOption) (*clientv3.GetResponse, error) {
	return cli.Get(ctx, "operation", opts...)
}

func putValueEtcd(value int64) {
	cli, ctx, cancel, err := etcdConnect()
	if err != nil {
		fmt.Println("Error opening ETCD connection")
	}

	keyToEtcd := fmt.Sprintf("operation_%02d", rand.Intn(1000))
	fmt.Printf("Key to ETCD : %s \n", keyToEtcd)
	fmt.Printf("Value to ETCD : %d \n", value)

	resPut, errPut := cli.Put(ctx, keyToEtcd, strconv.FormatInt(value, 10))
	if errPut != nil {
		fmt.Println(errPut)
	} else {
		fmt.Printf("Put on ETCD : %s \n", resPut.OpResponse().Put())
	}

	closeEtcd(cli, cancel)
}

func etcdConnect() (*clientv3.Client, context.Context, context.CancelFunc, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Etcd Client connected")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	return cli, ctx, cancel, err
}

func closeEtcd(cli *clientv3.Client, cancel context.CancelFunc) {
	cancel()
	errClose := cli.Close()
	if errClose != nil {
		fmt.Println("Error closing client")
	} else {
		fmt.Println("ETCD client disconnected")
	}
}

func subtract(w http.ResponseWriter, r *http.Request) {
	intValA, intValB, err := parseAndCheckValues(w, r)
	if err != nil {
		return
	}

	result := intValA - intValB
	putValueEtcd(result)

	fmt.Fprintf(w, "The subtraction result is: %d", result)
}

func multiply(w http.ResponseWriter, r *http.Request) {
	intValA, intValB, err := parseAndCheckValues(w, r)
	if err != nil {
		return
	}

	result := intValA * intValB
	putValueEtcd(result)

	fmt.Fprintf(w, "The multiplication result is: %d", result)
}

func divide(w http.ResponseWriter, r *http.Request) {
	intValA, intValB, err := parseAndCheckValues(w, r)
	if err != nil {
		return
	}

	result := intValA / intValB
	putValueEtcd(result)

	fmt.Fprintf(w, "The division result is: %d", result)
}

func sum(w http.ResponseWriter, r *http.Request) {
	intValA, intValB, err := parseAndCheckValues(w, r)
	if err != nil {
		return
	}

	result := intValA + intValB
	putValueEtcd(result)

	fmt.Fprintf(w, "The sum result is: %d", result)
}

func parseAndCheckValues(w http.ResponseWriter, r *http.Request) (int64, int64, error) {
	vars := mux.Vars(r)

	intValA, err := parseAndCheckIntValue(vars["valA"], w)
	if err != nil {
		return 0, 0, err
	}
	intValB, err := parseAndCheckIntValue(vars["valB"], w)
	if err != nil {
		return 0, 0, err
	}
	return intValA, intValB, nil
}

func parseAndCheckIntValue(val string, w http.ResponseWriter) (int64, error) {
	varA, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Parameter must '%s' be an Integer", val)
		return 0, err
	}
	return varA, err
}
