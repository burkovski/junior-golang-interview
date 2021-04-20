## Pointers:

1. Что такое указатель? Что выведет [этот код](https://codeshare.io/G7wxjY)? Почему? Какие изменения нужно внести, чтобы
   программа вывела 42 и 13?
   ```go
   package main
   
   import (
       "fmt"
   )
   
   func main() {
       v := 42
       fmt.Println(v)
       assignToPointer(&v)
       fmt.Println(v)
   }
   
   func assignToPointer(dst *int) {
       v := 13
       dst = &v
   }
   ```

## Slices:

1. Что такое слайс? В чем его отличие от массива? Что выведет [этот код](https://codeshare.io/5RPz14)? Почему?
   ```go
   package main
   
   import "fmt"
   
   func main() {
       s := []int{1, 2, 3, 4, 5}
       s1 := s[:]
       s2 := s[1:3]
   
       fmt.Printf("%#v\n", s)
       s1[0] = 42
       s2[0] = 42
       fmt.Printf("%#v\n", s)
   }
   ```

2. Как работает append? Что выведет [этот код](https://codeshare.io/2KAop8)? Почему?
   ```go
   package main
   
   import "fmt"
   
   func main() {
      s := []int{1, 2, 3, 4, 5}
      s1 := s[1:3]
   
      fmt.Printf("%#v\n", s)
      appendToSlice(s)
      appendToSlice(s1)
      fmt.Printf("%#v\n", s)
   }
   
   func appendToSlice(dst []int) {
      dst = append(dst, 42)
   }
   ```

3. Какое у slice zero value? Какие операции над ним возможны? Что выведет [этот код](https://codeshare.io/aygqyv)?
   Почему?
   ```go
   package main
   
   import "fmt"
   
   func main() {
      var s []string
      s = append(s, s...)
      s = append(s, "foobar")
      s = append(s, s...)
   
      for i := range s {
         fmt.Printf("%#v\n", s[i])
      }
   }
   ```

4. Что выведет [этот код](https://codeshare.io/5zZgA7)? Почему? Как его исправить?
   ```go
   package main
   
   import "fmt"
   
   type User struct {
       ID         int
       IsVerified bool
   }
   
   func main() {
       users := []User{{1, true}, {2, false}, {3, true}}
       verified := filterVerifiedUsers(users)
       for _, user := range verified {
           fmt.Println(user.ID)
       }
   }
   
   func filterVerifiedUsers(users []User) []*User {
       var verified []*User
       for _, user := range users {
           if user.IsVerified {
               verified = append(verified, &user)
           }
       }
       return verified
   }
   ```

## Maps:

1. Что такое map? Как устроен этот тип? Каков порядок перебора map? Что выведет [этот код](https://codeshare.io/GkknZO)?
   Почему? Как его исправить?
   ```go
   package main
   
   import "fmt"
   
   type User struct {
       ID int
   }
   
   func main() {
       var users map[int]User
       for i := 0; i < 1000; i++ {
           users[i] = User{i}
       }
       fmt.Printf("%#v\n", users)
   }
   ```

2. Можно ли конкурентно писать в map? Что выведет [этот код](https://codeshare.io/5v1xDm)? Почему? Как его исправить?
   ```go
   package main
   
   import (
       "fmt"
       "sync"
   )
   
   type User struct {
       ID int
   }
   
   func main() {
       users := make(map[int]User)
       wg := sync.WaitGroup{}
       for i := 0; i < 1000; i++ {
           wg.Add(1)
           go func(i int) {
               defer wg.Done()
               users[i] = User{i}
           }(i)
       }
       wg.Wait()
       fmt.Printf("%#v\n", users)
   }
   ```

## Channels:

1. Что такое канал? Для чего применятся? Какие бывают каналы? Что выведет [этот код](https://codeshare.io/aJ0PXR)?
   Почему? Как его исправить?
   ```go
   package main
   
   func main() {
           var c chan string
           c <- "let's get started"
   }
   ```

2. Что выведет [этот код](https://codeshare.io/G66mbO)? Почему? Как его исправить?
   ```go
   package main
   
   import "fmt"
   
   func main() {
       var c chan string
       fmt.Println(<-c)
   }
   ```

3. Что выведет [этот код](https://codeshare.io/GAkgbE)? Почему? Как его исправить?
   ```go
   package main
   
   import "fmt"
   
   func main() {
       var c = make(chan int, 100)
       for i := 0; i < 10; i++ {
           go func() {
               for j := 0; j < 10; j++ {
                   c <- j
               }
               close(c)
           }()
       }
       for i := range c {
           fmt.Println(i)
       }
   }
   ```   

4. Что выведет [этот код](https://codeshare.io/5PvX7w)? Почему? Как его исправить?
   ```go
   package main
   
   import "fmt"
   
   func main() {
       c := make(chan int, 3)
       c <- 1
       c <- 2
       c <- 3
       close(c)
       for i := 0; i < 4; i++ {
           fmt.Printf("%d ", <-c)
       }
   }
   ```

## Goroutines:

1. Почему конкурентность != параллелизм? Какой вид конкурентности используется в Golang? Что такое goroutine? Чем
   отличается от системного потока? Кто и когда переключает горутины?

2. Что выведет [этот код](https://codeshare.io/adKABM)? Почему? Как его исправить?
   ```go
   package main
   
   import (
       "fmt"
       "sync"
   )
   
   func main() {
       ch := make(chan int)
       wg := &sync.WaitGroup{}
       for i := 0; i < 3; i++ {
           wg.Add(1)
           go func(i int) {
               ch <- (i + 1) * 2
               wg.Done()
           }(i)
       }
       fmt.Println(<-ch)
       wg.Wait()
   }
   ```

3. Какие есть примитивы синхронизации? Что выведет [этот код](https://codeshare.io/2j38BP)? Почему? Как его исправить?
   ```go
   package main
   
   import "fmt"
   
   func main() {
       var counter int
       for i := 0; i < 1000; i++ {
           go func() {
               counter++
           }()
       }
       fmt.Println(counter)
   }
   ```

4. Можно ли реализовать sync.Mutex и sync.WaitGroup на каналах? Как?
5. Что ты использовал из пакета sync(кроме Mutex и WaitGroup)?
6. За сколько примерно выполнится [приложение](https://codeshare.io/5w7Zm9) — за 3 секунды или за 6? Что нужно изменить,
   чтобы код работал за 3 секунды?
   ```go
   package main
   
   import (
       "fmt"
       "time"
   )
   
   func main() {
       timeStart := time.Now()
   
       _, _ = <-worker(), <-worker()
   
       elapsed := time.Since(timeStart).Seconds()
       fmt.Println(int(elapsed))
   }
   
   func worker() chan int {
       ch := make(chan int)
       go func() {
           time.Sleep(3 * time.Second)
           ch <- 42
       }()
       return ch
   }
   ```