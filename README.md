
# 🚀 How to Run the Project

### Step 1: Install Dependencies

Ensure you have:
- **Go** (version 1.18 or higher).
- A modern browser to access the user interface.

### Step 2: Start the Server

Run the following command in your terminal:

```bash
go run main.go
```

### Step 3: Access the Interface

Open your browser and navigate to:

```
http://localhost:8080/
```

---

### 🌟 Features and How It Works

1. **Add Items to Cache**  
   Enter a key and value, then click `Add to Cache` to store the data in the LRU Cache.

2. **Retrieve Items**  
   Input a key and click `Get from Cache` to fetch a value. A success message or a `Cache miss` will be displayed.

3. **Automatic Removal**  
   Observe the least recently used items being removed (with smooth animations!) when the cache reaches its capacity.

4. **Deleted Items History**  
   All removed items are logged in a history table, including their key, value, and timestamp of removal.

---

### 👥 Meet the Team

- **Sanzhik** — Lead Developer  
- **Ilyas** — Meme Specialist  
- **Dancho** — Style and UX Designer  

![LRU Cache Crew](static/crew.png)

---

### 📜 License

This project is built for educational purposes only. Use it to learn algorithms, practice coding, and improve your skills!

---

### **How to Push the Updated README to GitHub**

1. Save this updated `README.md` file in your project directory.  
2. Ensure all required assets (e.g., `crew.png`, `screenshot.png`) are properly placed in the `static` folder.  
3. Push the changes to GitHub:

```bash
git add README.md static/crew.png
git commit -m "Updated README with improved layout and meme"
git push origin main
```

4. Verify the changes  
   Open your GitHub repository in a browser and check the updated `README.md` file to ensure everything appears as intended.
