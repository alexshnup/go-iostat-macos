Overview
--------

Extended iostat for MacOS based github.com/lufia/iostat with CGO

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
  
  Device:            MB_read/s MB_wrtn/s  #_read/s  #_wrtn/s  T_read/ms  T_wrtn/ms  R_lat(ms)  W_lat(ms)  #_r_err    #_w_err    #_r_retr    #_w_retr
  
  disk0                38.32     21.08   9809.04   5084.84      3.09     15.62       0.00       0.00       0.00       0.00        0.00        0.00
  disk2                 0.00      0.00      0.00      0.00      0.00      0.00       0.00       0.00       0.00       0.00        0.00        0.00




 $ ./iostat 1 2
 OS kernel version is darwin
  
 Device:            MB_read/s MB_wrtn/s  #_read/s  #_wrtn/s  T_read/ms  T_wrtn/ms  R_lat(ms)  W_lat(ms)  #_r_err    #_w_err    #_r_retr    #_w_retr
  
 disk0                38.32     21.08   9809.04   5084.84      3.09     15.62       0.00       0.00       0.00       0.00        0.00        0.00
 disk2                 0.00      0.00      0.00      0.00      0.00      0.00       0.00       0.00       0.00       0.00        0.00        0.00

 disk0                40.54     33.42  10377.32   7290.31      1.92     11.28       0.00       0.00       0.00       0.00        0.00        0.00
 disk2                 0.00      0.00      0.00      0.00      0.00      0.00       0.00       0.00       0.00       0.00        0.00        0.00



 $ ./iostat 1 2 disk0
 OS kernel version is darwin
  
 Device:            MB_read/s MB_wrtn/s  #_read/s  #_wrtn/s  T_read/ms  T_wrtn/ms  R_lat(ms)  W_lat(ms)  #_r_err    #_w_err    #_r_retr    #_w_retr
  
 disk0                38.32     21.08   9809.04   5084.84      3.09     15.62       0.00       0.00       0.00       0.00        0.00        0.00
 disk0                40.54     33.42  10377.32   7290.31      1.92     11.28       0.00       0.00       0.00       0.00        0.00        0.00

 $ ./iostat disk0
 OS kernel version is darwin
 
 Device:            MB_read/s MB_wrtn/s  #_read/s  #_wrtn/s  T_read/ms  T_wrtn/ms  R_lat(ms)  W_lat(ms)  #_r_err    #_w_err    #_r_retr    #_w_retr
 
 disk0                38.32     21.08   9809.04   5084.84      3.09     15.62       0.00       0.00       0.00       0.00        0.00        0.00
 disk0                40.54     33.42  10377.32   7290.31      1.92     11.28       0.00       0.00       0.00       0.00        0.00        0.00
 disk0                38.32     21.08   9809.04   5084.84      3.09     15.62       0.00       0.00       0.00       0.00        0.00        0.00
 disk0                40.54     33.42  10377.32   7290.31      1.92     11.28       0.00       0.00       0.00       0.00        0.00        0.00



