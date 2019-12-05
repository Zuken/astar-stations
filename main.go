package main

import (
   "github.com/jupp0r/go-priority-queue"
)

type Station struct {
   Code string
   Connections []*Connection
   Visited bool
}

func (source *Station) ConnectBidirect(target *Station, distance int) *Station {
   source.Connections = append(source.Connections, &Connection{distance, "both", source, target})
   target.Connections = append(target.Connections, &Connection{distance, "both", target, source})

   return target
}
func (source *Station) ConnectTo(target *Station, distance int) *Station {
   connection := Connection{distance, "both", source, target}
   source.Connections = append(source.Connections, &connection)

   return target
}

func (s *Station) GetConnection(target *Station) *Connection {
   for _,connection := range s.Connections {
      if connection.To == target {
         return connection
      }
   }
   return nil
}
func (s *Station) HasUnvisitedNeighbours() bool {
   for _, n := range s.Neighbours() {
      if n.Visited == false {
         return true
      }
   }
   return false
}
func (s *Station) Neighbours() []*Station {
   var items []*Station
   for _, connection := range s.Connections {
      if connection.To != s {
         items = append(items, connection.To)
      }
   }
   return items
}

type Connection struct {
   Length int
   Direction string
   From *Station
   To *Station
}

func cost(nextStation *Station, currentStation *Station) int {
   conn :=  currentStation.GetConnection(nextStation)
   if conn != nil {
      return conn.Length
   }
   return 0
}
func heuristic(finishStation *Station, currentStation *Station) int {
   //conn :=  currentStation.GetConnection(finishStation)
   //if conn != nil {
   // return conn.Length
   //}
   return 1
}

func main() {
   s1 := Station{Code: "1"}
   s2 := Station{Code: "2"}
   s3 := Station{Code: "3"}
   s4 := Station{Code: "4"}
   s5 := Station{Code: "5"}
   s6 := Station{Code: "6"}
   s7 := Station{Code: "7"}
   s8 := Station{Code: "8"}

   s8.ConnectBidirect(&s1, 5).
      ConnectBidirect(&s2, 5).
      ConnectBidirect(&s3, 17).
      ConnectBidirect(&s5, 7)

   s2.ConnectBidirect(&s4, 3).
      //ConnectTo(&s6, 3).
      ConnectBidirect(&s6, 3).
      ConnectBidirect(&s3, 1)

   s2.ConnectBidirect(&s7, 3)

   queue := pq.New()

   fromStation := &s5
   finishStation := &s1

   queue.Insert(fromStation, 0)
   cameFrom := map[*Station]*Station{};
   costCurrent := map[*Station]int{};
   costCurrent[fromStation] = 0
   for queue.Len()>0 {
      item, _ := queue.Pop()

      currentStation := item.(*Station)
      currentStation.Visited = true

      if currentStation == finishStation {
         break
      }

      for _, nextStation := range currentStation.Neighbours() {
         newCost := costCurrent[currentStation] + cost(currentStation, nextStation)
         if (cameFrom[nextStation] == nil || newCost < costCurrent[nextStation]) {
            costCurrent[nextStation] = newCost
            priority := newCost + heuristic(finishStation, nextStation)
            queue.Insert(nextStation, float64(priority))
            cameFrom[nextStation] = currentStation
         }
      }

   }

   c := finishStation
   path := []*Station{finishStation}
   for c != fromStation {
      c = cameFrom[c]
      path = append(path, c)
   }
   for _, s := range path {
      println(s.Code)
   }

}
