
func TestAuthentication(t *testing.T) {
    // Create an initialized set of authentication options
    // based on available OS_* environment variables.
    ao, err := utils.AuthOptions()
    if err != nil {
        t.Error(err)
        return
    }

    // Attempt to authenticate with them.
    r, err := identity.Authenticate(ao)
    if err != nil {
        t.Error(err)
        return
    }

    // We're authenticated; now let's grab our authen-
    // tication token.
    tok, err := identity.GetToken(r)
    if err != nil {
        t.Error(err)
        return
    }

    // Authentication tokens have a variety of fields which
    // might be of some interest.  Let's print a few of them out.

    table := map[string]string{
        "ID":      tok.Id,
        "Expires": tok.Expires,
    }

    for attr, val := range table {
        fmt.Printf("Your token's %s is %s\n", attr, val)
    }

    // With each authentication, you receive a master dir-
    // ectory of all the services your account can access.
    // This "service catalog", as OpenStack calls it, pro-
    // vides you the means to exploit other OpenStack ser-
    // vices.
    sc, err := identity.GetServiceCatalog(r)
    if err != nil {
        t.Error(err)
        return
    }

    // Prepare our elastic tabstopped writer for our table.
    w := new(tabwriter.Writer)
    w.Init(os.Stdout, 2, 8, 2, ' ', 0)

    // Different providers will provide different services.
    // Let's print them in summary.
    ces, err := sc.CatalogEntries()
    fmt.Println("Service Catalog Summary:")
    fmt.Fprintln(w, "Name\tType\t")
    for _, ce := range ces {
        fmt.Fprintf(w, "%s\t%s\t\n", ce.Name, ce.Type)
    }
    w.Flush()

    // Now let's print them in greater detail.
    for _, ce := range ces {
        fmt.Printf("Endpoints for %s/%s\n", ce.Name, ce.Type)
        fmt.Fprintln(w, "Version\tRegion\tTenant\tPublic URL\tInternal URL\t")
        for _, ep := range ce.Endpoints {
            fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t\n", ep.VersionId, ep.Region, ep.TenantId, ep.PublicURL, ep.InternalURL)
        }
        w.Flush()
    }
}
