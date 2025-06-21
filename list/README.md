# go-blog

你可以使用以下内容填充 [README.md](file:///Users/Ypuyu/Desktop/workspace/Github/AcWing/README.md) 文件，以清晰地描述你实现的**虚拟头结点方式的双向链表（Doubly Linked List with Dummy Head）**：

---

# 双向链表（Doubly Linked List）实现说明

本项目基于 Go 语言实现了一个**带有虚拟头结点和虚拟尾结点的双向链表（Doubly Linked List）**，旨在简化边界条件处理并提高代码可读性与健壮性。

- 使用**虚拟头结点（dummyHead）和虚拟尾结点（dummyTail）**，避免空指针判断。
- 所有插入、删除操作均统一处理，无需额外判断是否为头尾节点。
- 提供正向和反向遍历接口，返回切片便于测试和调试。

---

## 📷 示例输出

运行程序将依次输出如下结果：
```bash
[4 5 6 8]
[0 1 2 4 5 6 8]
[0 1 2 3 4 5 6 8]
[0 1 2 3 4 5 6 7 8]
[1 2 3 4 5 6 7]
[1 2 3 5 6 7]
[1 2 3 5 6 7]
[7 6 5 3 2 1]
```

---

## 🚀 后续拓展建议

- 支持并发安全版本（sync.Mutex）
- 支持迭代器模式

---