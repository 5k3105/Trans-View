package main

/*
#cgo CFLAGS: -pipe -fno-keep-inline-dllexport -O2 -Wall -Wextra -DUNICODE -DQT_NO_DEBUG -DQT_GUI_LIB -DQT_CORE_LIB -DQT_NEEDS_QMAIN
#cgo CXXFLAGS: -pipe -fno-keep-inline-dllexport -O2 -std=gnu++11 -frtti -Wall -Wextra -fexceptions -mthreads -DUNICODE -DQT_NO_DEBUG -DQT_GUI_LIB -DQT_CORE_LIB -DQT_NEEDS_QMAIN
#cgo CXXFLAGS: -I../../5k3105 -I. -IC:/Qt/Qt5.8.0/5.8/mingw53_32/include -IC:/Qt/Qt5.8.0/5.8/mingw53_32/include/QtGui -IC:/Qt/Qt5.8.0/5.8/mingw53_32/include/QtANGLE -IC:/Qt/Qt5.8.0/5.8/mingw53_32/include/QtCore -Irelease -IC:/Qt/Qt5.8.0/5.8/mingw53_32/mkspecs/win32-g++
#cgo LDFLAGS:        -Wl,-s -Wl,-subsystem,windows -mthreads
#cgo LDFLAGS:        -lmingw32 -LC:/Qt/Qt5.8.0/5.8/mingw53_32/lib -lqtmain -LC:/utils/my_sql/my_sql/lib -LC:/utils/postgresql/pgsql/lib -lshell32 -lQt5Core -lQt5Gui -lQt5Core
#cgo LDFLAGS: -Wl,--allow-multiple-definition
*/
import "C"
