const taskInput = document.getElementById("taskInput");
const addTaskBtn = document.getElementById("addTaskBtn");
const taskList = document.getElementById("taskList");
const taskCount = document.getElementById("taskCount");

const loadTasks = () => {
  const tasks = JSON.parse(localStorage.getItem("tasks")) || [];
  tasks.forEach(({ text, completed }) => addTask(text, completed));
  updateTaskCount();
};
const saveTasks = () => {
  const tasks = [...taskList.children].map((task) => ({
    text: task.querySelector(".task-text").textContent,
    completed: task.querySelector("input[type='checkbox']").checked,
  }));
  localStorage.setItem("tasks", JSON.stringify(tasks));
};
// Add task
const addTask = (taskText, completed = false) => {
  if (taskText.trim() === "") {
    alert("Task cannot be empty!");
    return;
  }

  const li = document.createElement("li");
  if (completed) li.classList.add("completed");
  li.innerHTML = `
    <input type="checkbox" ${completed ? "checked" : ""}>
    <span class="task-text">${taskText}</span>
    <div class="task-buttons">
      <button class="edit-btn">Edit</button>
      <button class="delete-btn">Delete</button>
    </div>
  `;
  const checkbox = li.querySelector("input[type='checkbox']");
  checkbox.addEventListener("change", () => {
    li.classList.toggle("completed");
    saveTasks();
    updateTaskCount();
  });
  // edit task
  const editBtn = li.querySelector(".edit-btn");
  editBtn.addEventListener("click", () => {
    const taskTextSpan = li.querySelector(".task-text");
    const input = document.createElement("input");
    input.type = "text";
    input.value = taskTextSpan.textContent;

    taskTextSpan.replaceWith(input);
    editBtn.textContent = "Save";

    input.focus();

    // save changes
    const saveChanges = () => {
      if (input.value.trim()) {
        taskTextSpan.textContent = input.value.trim();
        input.replaceWith(taskTextSpan);
        editBtn.textContent = "Edit";
        saveTasks();
      } else {
        alert("Task cannot be empty!");
        input.focus();
      }
    };

    editBtn.addEventListener("click", saveChanges, { once: true });
  });

  //deletion
  li.querySelector(".delete-btn").addEventListener("click", () => {
    li.remove();
    saveTasks();
    updateTaskCount();
  });

  taskList.appendChild(li);
  taskInput.value = "";
  updateTaskCount();
  saveTasks();
};
// pending tasks count
const updateTaskCount = () => {
  const pendingTasks = [...taskList.children].filter(
    (task) => !task.querySelector("input[type='checkbox']").checked
  ).length;
  taskCount.textContent = `Pending tasks: ${pendingTasks}`;
};

addTaskBtn.addEventListener("click", () => addTask(taskInput.value));

loadTasks();
