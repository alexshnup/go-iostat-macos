Overview
--------

Extended iostat for MacOS based on package https://github.com/lufia/iostat with CGO Apple API calls

Available args:
iostat [delay] [count] [disk]


Building
--------

Just type::

 $ go mod tidy
 $ go build .


Example
--------

Just type::
  
  $ ./iostat
  OS kernel version is darwin
  
  Device:            MB_read/s MB_wrtn/s  #_read/s  #_wrtn/s  T_read/ms  T_wrtn/ms   utils
  
  disk0                64.00      0.01   1054.53      2.00      0.74       0.01      45.13
  disk2                 0.01     79.43      1.00    635.45      0.04      18.08      95.27
  
  
  $ ./iostat 1 2
  OS kernel version is darwin
  
  Device:            MB_read/s MB_wrtn/s  #_read/s  #_wrtn/s  T_read/ms  T_wrtn/ms   utils
  
  disk0                64.11      0.04   2103.34      4.99      1.04       0.00      53.34
  disk2                 0.01     83.31      1.00    666.45      0.08      19.15      95.55
  
  disk0               128.86      0.01   8025.24      2.00      3.41       0.00      80.01
  disk2                 0.00     86.66      0.00    693.31      0.00      19.11      95.54
  
  
  
 $ ./iostat 1 2 disk0
 OS kernel version is darwin
 
 Device:            MB_read/s MB_wrtn/s  #_read/s  #_wrtn/s  T_read/ms  T_wrtn/ms   utils
 
 disk0                95.39      0.10   4873.92      4.00      3.23       0.03      78.69
 disk0                77.28      0.00   3030.32      0.00      2.74       0.00      75.22
 
 $ ./iostat disk0
 OS kernel version is darwin
 
 Device:            MB_read/s MB_wrtn/s  #_read/s  #_wrtn/s  T_read/ms  T_wrtn/ms   utils
 
 disk0               128.19      0.19   4269.64     12.97      2.23       0.00      72.40



