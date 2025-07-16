import { useEffect, useState } from "react"
import axios from "axios"

type Todo = {
  id: number
  title: string
  completed: boolean
}

function App() {
  const API_BASE = process.env.REACT_APP_API_BASE_URL || "http://localhost:8080"
  const [todos, setTodos] = useState<Todo[]>([])
  const [input, setInput] = useState("")

  useEffect(() => {
    axios.get<Todo[]>(`${API_BASE}/api/todos`).then(res => {
      setTodos(res.data)
    })
  }, [])

  const addTodo = () => {
    if (!input.trim()) return
    axios.post<Todo>(`${API_BASE}/api/todos`, { title: input })
      .then(res => {
        setTodos([...todos, res.data])
        setInput("")
      })
  }

  const deleteTodo = (id: number) => {
    axios.delete(`${API_BASE}/api/todos`, { params: { id } })
      .then(() => {
        setTodos(todos.filter(todo => todo.id !== id));
      });
  };

  const toggleCompleted = async (id: number, completed: boolean) => {
    try {
      const res = await axios.patch<Todo>(`${API_BASE}/api/todos?id=${id}`, {
        completed: !completed
      })
      const updated: Todo = res.data
      setTodos(prev =>
        prev.map(todo =>
          todo.id === updated.id ? updated : todo
        )
      )
    } catch (err) {
      console.error("PATCHエラー:", err)
    }
  }

  return (
    <div className="main">
      <h1>ToDo アプリ (Go + React)</h1>
      <p>Goで作成されたREST APIからToDoリストを取得し、Reactで表示・追加・削除できるシンプルなサンプルアプリです。</p>
      <ul className="todo-list">
        {todos.map((todo: any) => (
          <li key={todo.id}>
            <input
              type="checkbox"
              checked={todo.completed}
              onChange={() => toggleCompleted(todo.id, todo.completed)}
            />
            <span style={{
              textDecoration: todo.completed ? "line-through" : "none",
              marginLeft: "8px"
            }}>{todo.title}</span>
            <button onClick={() => deleteTodo(todo.id)}>削除</button>
          </li>
        ))}
      </ul>
      <input
        value={input}
        onChange={e => setInput(e.target.value)}
        placeholder="新しいToDoを入力"
      />
      <button onClick={addTodo}>追加</button>
    </div>
  )
}

export default App
