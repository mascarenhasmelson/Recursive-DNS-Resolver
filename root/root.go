package root


// https://www.iana.org/domains/root/servers -IPv4
var(
  SERVERS=[]string{"202.12.27.33","199.7.83.42","198.41.0.4","193.0.14.129","192.58.128.30",
    "192.36.148.17","198.97.190.53","192.112.36.4","192.5.5.241","192.203.230.10","199.7.91.13","192.33.4.12",
    "170.247.170.2",
  }
)

func ReturnIps()[]string{
  return SERVERS

}
