from const.scan_state import ScanSate
import sys


def update_state(stat: ScanSate, desc='') -> None:
    #  db update
    #  log into file
    pass


def die(stat=ScanSate.FAILED, desc='') -> None:
    # db update
    # log into file
    sys.exit(1)

