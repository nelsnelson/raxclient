package main

import (
    "fmt"
    "os"
    "github.com/rackspace/gophercloud"
    "github.com/rackspace/gophercloud/openstack"
    "github.com/rackspace/gophercloud/openstack/utils"
)

func main() {
    //auth_url := os.Getenv("OS_AUTH_URL")
    //username := os.Getenv("OS_USERNAME")
    //password := os.Getenv("OS_PASSWORD")
    //tenantid := os.Getenv("OS_TENANT_ID")
//
    //fmt.Println("OS_AUTH_URL:", auth_url)
    //fmt.Println("OS_USERNAME:", username)
    //fmt.Println("OS_PASSWORD:", password)
    //fmt.Println("OS_TENANT_ID:", tenantid)
//
    //ao := gophercloud.AuthOptions{
    //    IdentityEndpoint: "{auth_url}"
    //    Username: "{username}",
    //    Password: "{password}",
    //    TenantId: "{tenantid}",
    //    AllowReauth: true,
    //}
//
    //fmt.Println("AuthOptions.username: " + ao.Username)
//
    //// Attempt to authenticate with them.
    //r, err := identity.Authenticate(ao)
    //if err != nil {
    //    fmt.Println("Error authenticating: " + err)
    //    return
    //}

    fmt.Println("hello, world")
}

