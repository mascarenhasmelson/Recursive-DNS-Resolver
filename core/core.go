package core

import(
  "github.com/miekg/dns"
  "fmt"
  "net/http"
  "time"
  "math/rand"
  "github.com/mascarenhasmelson/Recursive-DNS-Resolver/root"
)

func Returnresults(query string, w http.ResponseWriter, flusher http.Flusher)([]dns.RR, error){

getIPS:=root.ReturnIps()
IP_Index:=rand.Intn(len(getIPS))
selectedIP := getIPS[IP_Index]

// fmt.Printf("Type of IP_Index: %T\n", getIPS)
client := new(dns.Client)

for{
    message := new(dns.Msg)
    message.SetQuestion(dns.Fqdn(query), dns.TypeA)
    fmt.Printf("Query %s about %s\n", selectedIP, query)
    serverAddress := fmt.Sprintf("%s:53",  selectedIP)
    fmt.Fprintf(w, "Query %s about %s\n", selectedIP, query)
    flusher.Flush()
    response, _, err := client.Exchange(message, serverAddress)
    if err != nil {
        fmt.Fprintf(w, "Error: %v\n", err)
          flusher.Flush()
          return nil,err
    }
    if len(response.Answer)>0{
        if cname, ans := response.Answer[0].(*dns.CNAME);ans {
				      return Returnresults(cname.Target,w,flusher)
			     }
         return response.Answer, nil
       }

       if len(response.Extra) == 0 && len(response.Ns) != 0 {
          ns := response.Ns[0].(*dns.NS)
			    nsIP, err := Returnresults(ns.Ns, w, flusher)
			    if err != nil {
				        return nil, fmt.Errorf("error")
			     }
      selectedIP = nsIP[0].(*dns.A).A.String()
      }else {
            NSfound := false
            for _, rr := range response.Extra {
				          record, ans := rr.(*dns.A)
				              if ans {
					                   selectedIP = record.A.String()
					                   NSfound = true
					                   break
				                     }
			                    }
      if !NSfound {
      				return nil, fmt.Errorf("break in the chain")
              }
    }
    time.Sleep(time.Millisecond * 500)
}
}
