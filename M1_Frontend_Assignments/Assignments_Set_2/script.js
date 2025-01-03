const amountInput = document.getElementById("amount");
const descriptionInput = document.getElementById("description");
const categorySelect = document.getElementById("category");
const addExpenseBtn = document.getElementById("addExpense");
const expenseTable = document.getElementById("expenseTable");
const summaryList = document.getElementById("summaryList");

let expenses = JSON.parse(localStorage.getItem("expenses")) || [];

const loadExpenses = () => {
    expenseTable.innerHTML = "";
    expenses.forEach((expense, index) => {
        const row = document.createElement("tr");
        row.innerHTML = `
      <td>Rs ${expense.amount}</td>
      <td>${expense.description}</td>
      <td>${expense.category}</td>
      <td><button class="delete-btn" onclick="deleteExpense(${index})">Delete</button></td>
    `;
        expenseTable.appendChild(row);
    });
};
//add expense
const addExpense = (amount, description, category) => {
    if (!amount || !description) {
        alert("Please fill out all fields!");
        return;
    }
    const expense = { amount: parseFloat(amount), description, category };
    expenses.push(expense);
    localStorage.setItem("expenses", JSON.stringify(expenses));
    loadExpenses();
    updateSummary();
};
const updateSummary = () => {
    const summary = {};
    //calculating category summary
    expenses.forEach(expense => {
        if (summary[expense.category]) {
            summary[expense.category] += expense.amount;
        } else {
            summary[expense.category] = expense.amount;
        }
    });
    summaryList.innerHTML = "";
    for (const [category, total] of Object.entries(summary)) {
        const li = document.createElement("li");
        li.textContent = `${category}: Rs.${total}`;
        summaryList.appendChild(li);
    }
};
// delete a expense
const deleteExpense = (index) => {
    expenses.splice(index, 1);
    localStorage.setItem("expenses", JSON.stringify(expenses));
    loadExpenses();
    updateSummary();
};

addExpenseBtn.addEventListener("click", () => {
    const amount = amountInput.value;
    const description = descriptionInput.value;
    const category = categorySelect.value;

    addExpense(amount, description, category);

    amountInput.value = "";
    descriptionInput.value = "";
});

loadExpenses();
updateSummary();
