package core

import (
    )

type UbuntuKernelsDB struct {
    indexByKernelVersion map[string]string
    indexBySeriesName    map[string]string
}

func NewUbuntuKernelsDB() *UbuntuKernelsDB {
    return &UbuntuKernelsDB{}
}

func (self *UbuntuKernelsDB) IndexByKernelVersion(version string) map[string]interface{} {
    if self.indexByKernelVersion == nil {
        self.indexByKernelVersion = make(map[string]string, len(DB))
        self.indexByKernelVersion["3.13.0"] = "14.04"
        self.indexByKernelVersion["3.11.0"] = "13.10"
        self.indexByKernelVersion["3.8.0"]  = "13.04"
        self.indexByKernelVersion["3.5.0"]  = "12.10"
        self.indexByKernelVersion["3.2.0"]  = "12.04"
        self.indexByKernelVersion["3.0.0"]  = "11.10"
        self.indexByKernelVersion["2.6.38"] = "11.04"
        self.indexByKernelVersion["2.6.35"] = "10.10"
        self.indexByKernelVersion["2.6.32"] = "10.04"
        self.indexByKernelVersion["2.6.31"] = "9.10"
        self.indexByKernelVersion["2.6.28"] = "9.04"
        self.indexByKernelVersion["2.6.27"] = "8.10"
        self.indexByKernelVersion["2.6.24"] = "8.04"
        self.indexByKernelVersion["2.6.22"] = "7.10"
        self.indexByKernelVersion["2.6.20"] = "7.04"
        self.indexByKernelVersion["2.6.15"] = "6.06"

    }

    return DB[self.indexByKernelVersion[version]]
}

func (self *UbuntuKernelsDB) IndexBySeriesName(seriesName string) map[string]interface{} {
    if self.indexBySeriesName == nil {
        self.indexBySeriesName = make(map[string]string, len(DB))

        self.indexBySeriesName["trusty"]   = "14.04"
        self.indexBySeriesName["saucy"]    = "13.10"
        self.indexBySeriesName["raring"]   = "13.04"
        self.indexBySeriesName["quantal"]  = "12.10"
        self.indexBySeriesName["precise"]  = "12.04"
        self.indexBySeriesName["oneiric"]  = "11.10"
        self.indexBySeriesName["natty"]    = "11.04"
        self.indexBySeriesName["maverick"] = "10.10"
        self.indexBySeriesName["lucid"]    = "10.04"
        self.indexBySeriesName["karmic"]   = "9.10"
        self.indexBySeriesName["jaunty"]   = "9.04"
        self.indexBySeriesName["intrepid"] = "8.10"
        self.indexBySeriesName["hardy"]    = "8.04"
        self.indexBySeriesName["gutsy"]    = "7.10"
        self.indexBySeriesName["feisty"]   = "7.04"
        self.indexBySeriesName["dapper"]   = "6.06"

    }

    return DB[self.indexBySeriesName[seriesName]]
}


var DB = map[string]map[string]interface{} {
        // Trusty (14.04)
        //
        "14.04" : map[string]interface{} {
            "development"    : true,
            "series_version" : "14.04",
            "kernel"         : "3.13.0",
            "name"           : "trusty",
            "supported"      : false,
            "packages"       : []string {
                "linux",
                "linux-meta",
                "linux-exynos5",
            },
            "dependent-packages" : map[string]map[string]string {
                "linux" : map[string]string {
                    "meta"   : "linux-meta",
                    "signed" : "linux-signed",
                },
            },
            "derivative-packages" : map[string][]string {
                "linux" : {},
            },
            "sha1" : "",
            "md5"  : "",
        },

        // Saucy (13.10)
        //
        "13.10" : map[string]interface{} {
            "development"    : false,
            "series_version" : "13.10",
            "kernel"         : "3.11.0",
            "name"           : "saucy",
            "supported"      : true,
            "packages"       : []string {
                "linux",
                "linux-meta",
                "linux-exynos5",
            },
            "dependent-packages" : map[string]map[string]string {
                "linux" : map[string]string {
                    "meta"   : "linux-meta",
                    "signed" : "linux-signed",
                },
                "linux-lowlatency" : map[string]string {
                    "meta" : "linux-meta-lowlatency",
                },
                "linux-ppc" : map[string]string {
                    "meta" : "linux-meta-ppc",
                },
            },
            "derivative-packages" : map[string][]string {
                "linux" : { "linux-lowlatency", "linux-ppc" },
            },
            "sha1" : "",
            "md5"  : "",
        },

        // Raring (13.04)
        //
        "13.04" : map[string]interface{} {
            "development" : false,
            "series_version" : "13.04",
            "kernel"    : "3.8.0",
            "name"      : "raring",
            "supported" : false,
            "packages"       : []string {
                "linux",
                "linux-meta",
                "linux-meta-ti-omap4",
            },
            "dependent-packages" : map[string]map[string]string {
                "linux" : map[string]string {
                    "meta"   : "linux-meta",
                    "signed" : "linux-signed",
                },
                "linux-lowlatency" : map[string]string {
                    "meta" : "linux-meta-lowlatency",
                },
                "linux-ppc" : map[string]string {
                    "meta" : "linux-meta-ppc",
                },
            },
            "derivative-packages" : map[string][]string {
                "linux" : { "linux-lowlatency", "linux-ppc" },
            },
            "sha1" : "",
            "md5" : "",
        },

        "12.10" : map[string]interface{} {
            "development" : false,
            "series_version" : "12.10",
            "kernel"    : "3.5.0",
            "name"      : "quantal",
            "supported" : true,
            "packages"       : []string {
                "linux",
                "linux-meta",
                "linux-ti-omap4",
                "linux-meta-ti-omap4",
                "linux-armadaxp",
                "linux-meta-armadaxp",
            },
            "dependent-packages" : map[string]map[string]string {
                "linux"             : map[string]string {
                    "meta"   : "linux-meta",
                    "lbm"    : "linux-backports-modules-3.5.0",
                    "signed" : "linux-signed",
                },
                "linux-ti-omap4"   : map[string]string {
                    "meta" : "linux-meta-ti-omap4",
                },
                "linux-armadaxp"   : map[string]string {
                    "meta" : "linux-meta-armadaxp",
                },
                "linux-lowlatency" : map[string]string {
                    "meta" : "linux-meta-lowlatency",
                },
            },
            "derivative-packages" : map[string][]string {
                "linux" : { "linux-ti-omap4", "linux-armadaxp", "linux-lowlatency" },
            },
            "sha1" : "",
            "md5" : "",
        },

        "12.04" : map[string]interface{} {
            "series_version" : "12.04",
            "kernel"    : "3.2.0",
            "name"      : "precise",
            "supported" : true,
            "packages"       : []string {
                "linux",
                "linux-meta",
                "linux-lts-quantal",
                "linux-meta-lts-quantal",
                "linux-signed-lts-quantal",
                "linux-lts-raring",
                "linux-meta-lts-raring",
                "linux-signed-lts-raring",
                "linux-lts-saucy",
                "linux-meta-lts-saucy",
                "linux-signed-lts-saucy",
                "linux-ti-omap4",
                "linux-meta-ti-omap4",
                "linux-armadaxp",
                "linux-meta-armadaxp",
            },
            "dependent-packages" : map[string]map[string]string {
                "linux" : map[string]string {
                    "meta" : "linux-meta",
                    "lbm"  : "linux-backports-modules-3.2.0",
                },
                "linux-lts-quantal" : map[string]string {
                    "meta" : "linux-meta-lts-quantal",
                    "signed" : "linux-signed-lts-quantal",
                },
                "linux-lts-raring" : map[string]string {
                    "meta" : "linux-meta-lts-raring",
                    "signed" : "linux-signed-lts-raring",
                },
                "linux-lts-saucy" : map[string]string {
                    "meta" : "linux-meta-lts-saucy",
                    "signed" : "linux-signed-lts-saucy",
                },
                "linux-ti-omap4" : map[string]string { "meta" : "linux-meta-ti-omap4" },
                "linux-armadaxp" : map[string]string { "meta" : "linux-meta-armadaxp" },
                "linux-lowlatency" : map[string]string { "meta" : "linux-meta-lowlatency" },
            },
            "derivative-packages" : map[string][]string {
                "linux" : { "linux-ti-omap4", "linux-armadaxp", "linux-lowlatency" },
            },
            "backport-packages"   : map[string][]string {
                "linux-lts-quantal" : { "linux", "12.10" },
                "linux-lts-raring"  : { "linux", "13.04" },
                "linux-lts-saucy"   : { "linux", "13.10" },
            },
            "sha1" : "",
            "md5" : "",
        },


        "11.10" : map[string]interface{} {
            "series_version" : "11.10",
            "kernel"    : "3.0.0",
            "name"      : "oneiric",
            "supported" : false,
            "packages"       : []string {
                "linux",
                "linux-meta",
                "linux-ti-omap4",
                "linux-meta-ti-omap4",
            },
            "dependent-packages" : map[string]map[string]string {
                "linux" : map[string]string {
                    "meta" : "linux-meta",
                    "lbm"  : "linux-backports-modules-3.0.0",
                },
                "linux-ti-omap4" : map[string]string { "meta" : "linux-meta-ti-omap4" },
            },
            "derivative-packages" : map[string][]string {
                "linux" : { "linux-ti-omap4" },
            },
            "sha1" : "",
            "md5" : "",
        },

        "11.04" : map[string]interface{} {
            "series_version" : "11.04",
            "kernel"    : "2.6.38",
            "name"      : "natty",
            "supported" : false,
            "packages"       : []string {
                "linux",
                "linux-meta",
                "linux-ti-omap4",
                "linux-meta-ti-omap4",
                "linux-backports-modules-2.6.38",
            },
            "dependent-packages" : map[string]map[string]string {
                "linux" : map[string]string {
                    "meta" : "linux-meta",
                    "lbm"  : "linux-backports-modules-2.6.38",
                },
                "linux-ti-omap4" : map[string]string { "meta" : "linux-meta-ti-omap4" },
            },
            "derivative-packages" : map[string][]string {
                "linux" : { "linux-ti-omap4" },
            },
            "sha1" : "0770b9d2483eaeee4b80aec1fd448586b882003e",
            "md5" : "cf0b587742611328f095da4b329e9fc7",
        },

        "10.10" : map[string]interface{} {
            "series_version" : "10.10",
            "kernel" : "2.6.35",
            "name" : "maverick",
            "supported" : false,
            "packages"       : []string {
                "linux",
                "linux-ti-omap4",
                "linux-mvl-dove",
                "linux-meta",
                "linux-ports-meta",
                "linux-meta-mvl-dove",
                "linux-meta-ti-omap4",
                "linux-backports-modules-2.6.35",
            },
            "dependent-packages" : map[string]map[string]string {
                "linux" : map[string]string {
                    "meta" : "linux-meta",
                    "ports-meta" : "linux-ports-meta",
                    "lbm": "linux-backports-modules-2.6.35",
                },
                "linux-ti-omap4" : map[string]string { "meta" : "linux-meta-ti-omap4" },
                "linux-mvl-dove" : map[string]string { "meta" :  "linux-meta-mvl-dove" },
            },
            "derivative-packages" : map[string][]string {
                "linux" : { "linux-ti-omap4", "linux-mvl-dove" },
            },
            "sha1" : "a2422f9281766ffe2f615903712819b9b0d9dd52",
            "md5" : "62001687bd94d1c0dd9a3654c64257d6",
        },

        "10.04" : map[string]interface{} {
            "series_version" : "10.04",
            "kernel" : "2.6.32",
            "name" : "lucid",
            "supported" : true,
            "packages"       : []string {
                "linux",
                "linux-fsl-imx51",
                "linux-mvl-dove",
                "linux-ec2",
                "linux-meta",
                "linux-ports-meta",
                "linux-meta-ec2",
                "linux-meta-mvl-dove",
                "linux-meta-fsl-imx51",
                "linux-backports-modules-2.6.32",
                "linux-lts-backport-oneiric",
                "linux-meta-lts-backport-oneiric",
                "linux-lts-backport-natty",
                "linux-meta-lts-backport-natty",
                "linux-lts-backport-maverick",
                "linux-meta-lts-backport-maverick",
            },
            "dependent-packages" : map[string]map[string]string {
                "linux" : map[string]string {
                    "meta" : "linux-meta",
                    "ports-meta" : "linux-ports-meta",
                    "lbm" : "linux-backports-modules-2.6.32",
                },
                "linux-ec2"       : map[string]string { "meta" : "linux-meta-ec2" },
                "linux-lts-backport-oneiric" : map[string]string {
                    "meta" : "linux-meta-lts-backport-oneiric",
                },
                "linux-lts-backport-natty" : map[string]string {
                    "meta" : "linux-meta-lts-backport-natty",
                },
                "linux-lts-backport-maverick" : map[string]string {
                    "meta" : "linux-meta-lts-backport-maverick",
                },
            },
            "derivative-packages" : map[string][]string {
                "linux" : { "linux-ec2" },
            },
            "backport-packages"   : map[string][]string {
                "linux-lts-backport-oneiric"  : { "linux", "11.10" },
                "linux-lts-backport-natty"    : { "linux", "11.04" },
                "linux-lts-backport-maverick" : { "linux", "10.10" },
            },
            "sha1" : "298cbfdb55fc64d1135f06b3bed3c8748123c183",
            "md5" : "4b1f6f6fac43a23e783079db589fc7e2",
        },

        "9.10" : map[string]interface{} {
            "series_version" : "9.10",
            "kernel" : "2.6.31",
            "name" : "karmic",
            "supported" : false,
            "packages"       : []string {
                "linux",
                "linux-fsl-imx51",
                "linux-mvl-dove",
                "linux-ec2",
                "linux-meta",
                "linux-ports-meta",
                "linux-meta-ec2",
                "linux-meta-mvl-dove",
                "linux-meta-fsl-imx51",
                "linux-backports-modules-2.6.31",
            },
            "sha1" : "6b19c2987b0e2d74dcdca2aadebd5081bc143b72",
            "md5" : "16c0355d3612806ef87addf7c9f8c9f9",
        },

        "9.04" : map[string]interface{} {
            "series_version" : "9.04",
            "kernel" : "2.6.28",
            "name" : "jaunty",
            "supported" : false,
            "packages"       : []string {},
            "sha1" : "92d6a293200566646fbb9215e0633b4b9312ad38",
            "md5" : "062c29b626a55f09a65532538a6184d4",
        },

        "8.10" : map[string]interface{} {
            "series_version" : "8.10",
            "kernel" : "2.6.27",
            "name" : "intrepid",
            "supported" : false,
            "packages"       : []string {},
        },

        "8.04" : map[string]interface{} {
            "series_version" : "8.04",
            "kernel" : "2.6.24",
            "name" : "hardy",
            "supported" : false,
            "packages"       : []string {
                "linux",
                "linux-meta",
                "linux-backports-modules-2.6.24",
                "linux-ubuntu-modules-2.6.24",
                "linux-restricted-modules-2.6.24",
            },
            "dependent-packages" : map[string]map[string]string {
                "linux" : map[string]string {
                    "meta" : "linux-meta",
                    "lbm" : "linux-backports-modules-2.6.24",
                    "lrm" : "linux-restricted-modules-2.6.24",
                    "lum" : "linux-ubuntu-modules-2.6.24",
                },
            },
            "sha1" : "ccccdc4759fd780a028000a1b7b15dbd9c60363b",
            "md5" : "e4aad2f8c445505cbbfa92864f5941ab",
        },

        "7.10" : map[string]interface{} {
            "series_version" : "7.10",
            "kernel" : "2.6.22",
            "name" : "gutsy",
            "supported" : false,
            "packages"       : []string {},
        },

        "7.04" : map[string]interface{} {
            "series_version" : "7.04",
            "kernel" : "2.6.20",
            "name" : "feisty",
            "supported" : false,
            "packages"       : []string {},
        },

        "6.06" : map[string]interface{} {
            "series_version" : "6.06",
            "kernel" : "2.6.15",
            "name" : "dapper",
            "supported" : false,
            "packages"       : []string {
                "linux-source-2.6.15",
                "linux-backports-modules-2.6.15",
                "linux-restricted-modules-2.6.15",
            },
        },

    }
