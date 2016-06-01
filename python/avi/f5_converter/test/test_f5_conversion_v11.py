import copy
import logging
import os
import unittest
import csv

import avi.f5_converter.f5_config_converter_v11 as f5_config_converter
import avi.f5_converter.f5_parser as f5_parser

gSAMPLE_CONFIG = None
log = logging.getLogger(__name__)


def setUpModule():
    cfg_file = open('bigip_v11.conf', 'r')
    cfg = cfg_file.read()
    global gSAMPLE_CONFIG
    gSAMPLE_CONFIG = cfg
    log.debug(' read cofig %s', gSAMPLE_CONFIG)


class Test(unittest.TestCase):

    LOG = logging.getLogger("converter-log")
    LOG.setLevel(logging.DEBUG)
    fh = logging.FileHandler(".." + os.path.sep + "output" + os.path.sep +
                             "converter.log", mode='a', encoding=None,
                             delay=False)
    fh.setLevel(logging.DEBUG)
    formatter = logging.Formatter(
        '%(asctime)s - %(name)s - %(levelname)s - %(message)s')
    fh.setFormatter(formatter)
    LOG.addHandler(fh)

    def test_config_conversion(self):
        f5_config_dict = f5_parser.parse_config(gSAMPLE_CONFIG, 11)
        assert f5_config_dict.get("virtual", None)
        assert f5_config_dict.get("monitor", None)
        assert f5_config_dict.get("pool", None)
        assert f5_config_dict.get("profile", None)
        assert f5_config_dict.get("node", None)
        f5_config_test = copy.deepcopy(f5_config_dict)
        avi_config_dict = f5_config_converter.convert_to_avi_dict(
            f5_config_dict, ".."+os.path.sep+"output", "disable",
            "certs", "api-upload")

        with open('../output/ConversionStatus.csv') as csvfile:
            reader = csv.DictReader(csvfile)
            for row in reader:
                assert row['Status'] != 'error'


if __name__ == "__main__":
    unittest.main()