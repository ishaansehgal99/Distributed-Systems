package config

import "os"

const PORT string = "6789"

const TCP_PORT string = "6790"

const BUFFER_SIZE int = 32768

const T_TIMEOUT = 2

const T_CLEANUP = 2

const PULSE_TIME = 500

const GOSSIP_FANOUT = 4

const RING_SIZE = 128

const SDFS_DIR = "sdfs/"

const FILE_PERM = os.FileMode(0777)

const NUM_REPLICAS = 4

const FILE_BUFFER_SIZE = 100 * 1024 * 1024
