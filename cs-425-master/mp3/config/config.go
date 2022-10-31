package config

import "os"

const PORT string = "6789"

const TCP_PORT string = "6790"

const MAPLEJUICE_PORT string = "6791"

const BUFFER_SIZE = 100 * 1024 * 1024

const T_TIMEOUT = 2

const T_CLEANUP = 2

const PULSE_TIME = 500

const GOSSIP_FANOUT = 4

const RING_SIZE = 128

const SDFS_DIR = "sdfs/"

const FILE_PERM = os.FileMode(0777)

const NUM_REPLICAS = 4

const FILE_BUFFER_SIZE = 100 * 1024 * 1024

const MAP_LINES = 100 // number of lines to read per map

const MAPLE_INPUT_FILE_PREFIX = "maple_input_"
