package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework ScriptingBridge
#import <Foundation/Foundation.h>
#import "chrome.h"
#import "safari.h"

@implementation NSString (NSString_Extended)

- (NSString *)urlencode {
    NSMutableString *output = [NSMutableString string];
    const unsigned char *source = (const unsigned char *)[self UTF8String];
    int sourceLen = strlen((const char *)source);
    for (int i = 0; i < sourceLen; ++i) {
        const unsigned char thisChar = source[i];
        if (thisChar == ' '){
            [output appendString:@"+"];
        } else if (thisChar == '.' || thisChar == '-' || thisChar == '_' || thisChar == '~' ||
                   (thisChar >= 'a' && thisChar <= 'z') ||
                   (thisChar >= 'A' && thisChar <= 'Z') ||
                   (thisChar >= '0' && thisChar <= '9')) {
            [output appendFormat:@"%c", thisChar];
        } else {
            [output appendFormat:@"%%%02X", thisChar];
        }
    }
    return output;
}
@end

void block(const char *cFocusUrl, const char *cUrls[10]) {
	@autoreleasepool {
		NSString *focusUrl = [NSString stringWithUTF8String:cFocusUrl];

		//mdls -name kMDItemCFBundleIdentifier -r /Applications/Google \Chrome.app/
		ChromeApplication *chrome = [SBApplication applicationWithBundleIdentifier:@"com.google.Chrome"];

		for (ChromeWindow *window in chrome.windows) {
			for (ChromeTab *tab in window.tabs) {
				for (int i = 0; i < 10; ++i) {
					const char *cUrl = cUrls[i];
					if (!cUrl) {
						break;
					}

					NSString *url = [NSString stringWithUTF8String:cUrl];

					if ([tab.URL rangeOfString:url].location != NSNotFound) {
						tab.URL = focusUrl;
					}
				}
			}
		}

		//mdls -name kMDItemCFBundleIdentifier -r /Applications/Safari.app/
		SafariApplication *safari = [SBApplication applicationWithBundleIdentifier:@"com.apple.Safari"];

		printf("Safari safari.windows\n");

		NSUInteger windowIndex = 0;
		for (SafariWindow *window in safari.windows) {
			NSUInteger tabIndex = 0;
			for (SafariTab *tab in window.tabs) {
				for (int i = 0; i < 10; ++i) {
					const char *cUrl = cUrls[i];
					if (!cUrl) {
						break;
					}

					NSString *url = [NSString stringWithUTF8String:cUrl];

					if ([tab.URL rangeOfString:url].location != NSNotFound) {
						NSString *urlEncodedString = [tab.URL urlencode];
						NSLog(@"%s", urlEncodedString.UTF8String);
						[[[[[safari windows] objectAtIndex:windowIndex] tabs] objectAtIndex:tabIndex] setURL:focusUrl];
					}
				}
				tabIndex++;
			}
			windowIndex++;
		}
    }
}
*/
import "C"
import (
	"unsafe"
	"time"
	"net/http"
	"fmt"
	"log"
	_ "net/http/pprof"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() (err error) {
	go func() {
		for true {
			a := GetCStringArray([]string{
				"ebay.com",
				"facebook.com",
				"youtube.com",
				"adadqwdyx",
				"adadqwdyx",
				"adadqwdyx",
				"adadqwdyx",
				"adadqwdyx",
				"adadqwdyx",
				"adadqwdyx",
			})

			C.block(C.CString("http://127.0.0.1:1234"), a)

			time.Sleep(100 * time.Millisecond)
		}
	}()

	http.HandleFunc("/", viewHandler)
	go http.ListenAndServe("127.0.0.1:8080", http.DefaultServeMux)
	err = http.ListenAndServe("127.0.0.1:1234", nil)

	return
}

func GetCStringArray(data []string) (**C.char) {
	cArray := C.malloc(C.size_t(10) * C.size_t(unsafe.Sizeof(uintptr(0))))

	// convert the C array to a Go Array so we can index it
	a := (*[1<<30 - 1]*C.char)(cArray)

	for idx, substring := range data {
		a[idx] = C.CString(substring)
	}

	return (**C.char)(cArray)
}

func ReleaseCStringArray(arr []*C.char) {
	for i := range arr {
		C.free(unsafe.Pointer(arr[i]))
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<!DOCTYPE html>
	<html>
	<head>
	<title>focus</title>
	<link href="https://fonts.googleapis.com/css?family=Akronim|Chelsea+Market|Damion|Kaushan+Script|Nunito:200i|Share:700i|Swanky+and+Moo+Moo" rel="stylesheet">
	<style type="text/css">
		html {
			height: 100%;
		}
	body {
		position: relative;
		height: 100%;
		margin: 0;
		background-repeat: no-repeat;
		background-attachment: fixed;


		/* Permalink - use to edit and share this gradient: http://colorzilla.com/gradient-editor/#1a237e+0,1e88e5+100 */
		background: #1a237e; /* Old browsers */
		background: -moz-linear-gradient(-45deg, #1a237e 0%, #1e88e5 100%); /* FF3.6-15 */
		background: -webkit-linear-gradient(-45deg, #1a237e 0%,#1e88e5 100%); /* Chrome10-25,Safari5.1-6 */
		background: linear-gradient(135deg, #1a237e 0%,#1e88e5 100%); /* W3C, IE10+, FF16+, Chrome26+, Opera12+, Safari7+ */
		filter: progid:DXImageTransform.Microsoft.gradient( startColorstr='#1a237e', endColorstr='#1e88e5',GradientType=1 ); /* IE6-9 fallback on horizontal gradient */
	}
	h1 {
		//font-family: 'Nunito', sans-serif;
		font-family: 'Kaushan Script', cursive;
		//font-family: 'Damion', cursive;
		//font-family: 'Share', cursive;
		//font-family: 'Chelsea Market', cursive;
		//font-family: 'Swanky and Moo Moo', cursive;
		//font-family: 'Akronim', cursive;
		margin: 0;
		color: #fff;
		font-size: 100px;
		position: absolute;
		top: 45%;
		left: 48%;
		margin-right: -50%;
		transform: translate(-50%, -50%) }
}

</style>
</head>
<body>
<h1>focus</h1>
<h1></h1>
</body>
</html>`)
}
