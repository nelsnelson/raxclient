
// waitForStatus waits up to 5 minutes (estimated)
// for a server to reach ACTIVE status.
func waitForStatus(ts *testState, s string) error {
    var (
        inProgress bool
        timeout    int
        err        error
    )

    for inProgress, timeout, err = countDown(ts, 300); inProgress; inProgress, timeout, err = countDown(ts, timeout) {
        if ts.createdServer.Status == s {
            fmt.Printf("Server created after %d seconds (approximately)\n", 300-timeout)
            break
        }
    }

    if err == errTimeout {
        fmt.Printf("Timeout -- I'm not waiting around.\n")
        err = nil
    }

    return err
}

// countDown implements a simple state machine to im-
// plement a crude countdown timer.
func countDown(ts *testState, timeout int) (bool, int, error) {
    if timeout & lt; 1 {
        return false, 0, errTimeout
    }
    time.Sleep(1 * time.Second)
    timeout--

    gr, err := servers.GetDetail(ts.client, ts.createdServer.Id)
    if err != nil {
        return false, timeout, err
    }

    ts.createdServer, err = servers.GetServer(gr)
    if err != nil {
        return false, timeout, err
    }

    return true, timeout, nil
}
