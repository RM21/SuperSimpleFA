

package main

import (
   "fmt"
   "net/http"
   "io"
   "os"
   "path/filepath"
   "bufio"
   "strings"
   "strconv"
   "log"
)

func main() {
   apiKey := "GN91GUX4C31M1UN3";
   avFunctionName := "TIME_SERIES_DAILY";
   avTemplateUrl := "https://www.alphavantage.co/query?function=%s&symbol=%s&apikey=%s&datatype=csv"
   tickerSymbol := "BRK.B"

   var formatted_url = fmt.Sprintf(avTemplateUrl, avFunctionName, tickerSymbol, apiKey)

   print(formatted_url + " \n")

   destinationFilepath, err := filepath.Rel(".", tickerSymbol + ".csv")

   if (err != nil) {

   }

   downloadFile(destinationFilepath, formatted_url)

   getMax(destinationFilepath);
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func downloadFile(filepath string, url string) error {

   // Create the file
   out, err := os.Create(filepath)
   if err != nil {
      return err
   }
   defer out.Close()

   // Get the data
   resp, err := http.Get(url)
   if err != nil {
      return err
   }
   defer resp.Body.Close()

   // Write the body to file
   _, err = io.Copy(out, resp.Body)
   if err != nil {
      return err
   }

   return nil
}

func getMax(filepath string) {
   file, err := os.Open(filepath)
   if err != nil {
      log.Fatal(err)
   }
   defer file.Close()

   scanner := bufio.NewScanner(file)

   maxDailyEntry := DailyEntry{}

   for scanner.Scan() {
      dailyEntryUnformatted := scanner.Text()

      dailyEntryArray := strings.Split(dailyEntryUnformatted, ",")

      if (len(dailyEntryArray) > 3) {
         highFloat64, err := strconv.ParseFloat(dailyEntryArray[2], 64)

         if (err != nil) {

         }

         dailyEntry := DailyEntry{
            Timestamp: dailyEntryArray[0],
            Open:      dailyEntryArray[1],
            High:      highFloat64,
            Close:     dailyEntryArray[3],
            Volume:    dailyEntryArray[4],
         }

         if (dailyEntry.High > maxDailyEntry.High) {
            maxDailyEntry = dailyEntry
         }
      }
   }

   println(maxDailyEntry.Timestamp)
   println(strconv.FormatFloat(maxDailyEntry.High, 'f', 6, 64))
}

type DailyEntry struct  {
   Timestamp string
   Open string
   High float64
   Close string
   Volume string
}