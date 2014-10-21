
func TestCreateDestroyServer(t *testing.T) {
    // Initialize our test state (ts) structure based on
    // available environment variable settings.
    ts, err := setupForCRUD()
    if err != nil {
        t.Error(err)
        return
    }

    // First, we need to pick a random name for it.
    ts.serverName = randomString("ACPTTEST", 16)
    fmt.Printf("Attempting to create server: %s\n", ts.serverName)

    // We have a valid endpoint to use thanks to
    // setupForCRUD(), above.  Let's instantiate a client
    // on that endpoint.
    ts.client = servers.NewClient(ts.ep, ts.a, ts.o)

    // Now we can request the server be created.
    // We assume ts.flavorId and ts.imageId were set up
    // via setupForCRUD().
    cr, err := servers.Create(ts.client, map[string]interface{}{
        "flavorRef": ts.flavorId,
        "imageRef":  ts.imageId,
        "name":      ts.serverName,
    })
    if err != nil {
        return err
    }

    // The server returned by the Create endpoint will
    // only give a small synopsis of the full set of
    // information.  To get more detailed info about
    // the server, you should call the GetDetail() method.
    // However, to do that, we need to know the server
    // ID.
    //
    // Unfortunately, the data returned above is opaque;
    // it's stored effectively as a raw blob of data.
    // We need to decode it into a more structured form,
    // to access fields like Id.
    ts.createdServer, err = servers.GetServer(cr)

    // Remember to delete the server when we're done using
    // it.
    //
    // We put this in a defer so that it gets executed even
    // in the face of errors or panics.
    defer func() {
        servers.Delete(ts.client, ts.createdServer.Id)
    }()

    // Wait for the server provisioning to complete.
    err = waitForStatus(ts, "ACTIVE")
    if err != nil {
        t.Error(err)
    }

    // At this point, you're free to use the server how-
    // ever you wish.
}
