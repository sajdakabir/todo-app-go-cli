It's a simple CLI Todo app in GO 


commands:
1: ./todo -list   ---> list the todo list
2: ./todo -add [task name] --->add the tase to the todo list
3: ./todo -complete=[index] ---> it's mark the task as a complete 
4: ./todo -del=[index]  ---> it's delete the task


examples :

~/Desktop/go_todo_app$ ./todo -list

╔═══╤════════════════╤═══════╤═════════════════════╤═════════════════════╗
║ # │      Task      │ Done? │           CreatedAt │         CompletedAt ║
╟━━━┼━━━━━━━━━━━━━━━━┼━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━╢
║ 1 │ ✅ Sample todo │ yes   │ 01 Jan 01 00:00 UTC │ 26 Jun 22 22:42 IST ║
║ 2 │ reading book   │ no    │ 26 Jun 22 23:56 IST │ 01 Jan 01 00:00 UTC ║
╟━━━┼━━━━━━━━━━━━━━━━┼━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━╢
║                        You have 1 pending todos                        ║
╚═══╧════════════════╧═══════╧═════════════════════╧═════════════════════╝


~/Desktop/go_todo_app$ ./todo -add buy milk
~/Desktop/go_todo_app$ ./todo -list

╔═══╤════════════════╤═══════╤═════════════════════╤═════════════════════╗
║ # │      Task      │ Done? │           CreatedAt │         CompletedAt ║
╟━━━┼━━━━━━━━━━━━━━━━┼━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━╢
║ 1 │ ✅ Sample todo │ yes   │ 01 Jan 01 00:00 UTC │ 26 Jun 22 22:42 IST ║
║ 2 │ reading book   │ no    │ 26 Jun 22 23:56 IST │ 01 Jan 01 00:00 UTC ║
║ 3 │ buy milk       │ no    │ 27 Jun 22 02:16 IST │ 01 Jan 01 00:00 UTC ║
╟━━━┼━━━━━━━━━━━━━━━━┼━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━╢
║                        You have 2 pending todos                        ║
╚═══╧════════════════╧═══════╧═════════════════════╧═════════════════════╝



~/Desktop/go_todo_app$ ./todo -complete=2
~/Desktop/go_todo_app$ ./todo -list
╔═══╤═════════════════╤═══════╤═════════════════════╤═════════════════════╗
║ # │      Task       │ Done? │           CreatedAt │         CompletedAt ║
╟━━━┼━━━━━━━━━━━━━━━━━┼━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━╢
║ 1 │ ✅ Sample todo  │ yes   │ 01 Jan 01 00:00 UTC │ 26 Jun 22 22:42 IST ║
║ 2 │ ✅ reading book │ yes   │ 26 Jun 22 23:56 IST │ 27 Jun 22 02:17 IST ║
║ 3 │ buy milk        │ no    │ 27 Jun 22 02:16 IST │ 01 Jan 01 00:00 UTC ║
╟━━━┼━━━━━━━━━━━━━━━━━┼━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━╢
║                        You have 1 pending todos                         ║
╚═══╧═════════════════╧═══════╧═════════════════════╧═════════════════════╝




~/Desktop/go_todo_app$ ./todo -del=2
~/Desktop/go_todo_app$ ./todo -list
╔═══╤════════════════╤═══════╤═════════════════════╤═════════════════════╗
║ # │      Task      │ Done? │           CreatedAt │         CompletedAt ║
╟━━━┼━━━━━━━━━━━━━━━━┼━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━╢
║ 1 │ ✅ Sample todo │ yes   │ 01 Jan 01 00:00 UTC │ 26 Jun 22 22:42 IST ║
║ 2 │ buy milk       │ no    │ 27 Jun 22 02:16 IST │ 01 Jan 01 00:00 UTC ║
╟━━━┼━━━━━━━━━━━━━━━━┼━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━┼━━━━━━━━━━━━━━━━━━━━━╢
║                        You have 1 pending todos                        ║
╚═══╧════════════════╧═══════╧═════════════════════╧═════════════════════╝