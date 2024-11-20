<template>
    <div class="flex h-screen">
        <!-- Sidebar -->
        <div class="w-64 bg-gray-800 text-white flex flex-col">
            <div class="p-4 text-2xl font-bold border-b border-gray-700">Task Manager</div>
            <ul class="flex-1 overflow-auto">
                <li @click="toggleDashboard" class="p-4 hover:bg-gray-700 cursor-pointer">
                    Dashboard
                </li>
            </ul>
            <div class="p-4 text-gray-400 text-sm border-t border-gray-700">
                © 2024 jirayu satsanatayat
            </div>
            <button @click="logout" class="bg-red-500 text-white px-4 py-2">Logout</button>

        </div>

        <!-- Main Content -->
        <div class="flex-1 p-6 bg-gray-50">
            <h1 class="text-3xl font-bold mb-6 text-gray-800">Task Board</h1>

            <!-- Board Layout -->
            <div class="grid grid-cols-3 gap-4">
                <div v-for="(column, index) in board" :key="index" class="bg-white rounded-lg shadow p-4">
                    <h2 class="text-xl font-semibold text-gray-700 mb-4">{{ column.name }}</h2>
                    <div class="space-y-3">
                        <div v-for="task in column.tasks" :key="task.id"
                            class="bg-gray-100 p-4 rounded-lg shadow hover:shadow-md transition-all">
                            <!-- Edit Mode -->
                            <div v-if="task.isEditing">
                                <input v-model="task.name" placeholder="Task Name"
                                    class="w-full border p-2 rounded-lg mb-2" />
                                <textarea v-model="task.description" placeholder="Description"
                                    class="w-full border p-2 rounded-lg mb-2"></textarea>
                                <input v-model="task.dueDate" type="date" class="w-full border p-2 rounded-lg mb-2" />
                                <select v-model="task.status" class="w-full border p-2 rounded-lg mb-2">
                                    <option value="todo">To Do</option>
                                    <option value="in-progress">In Progress</option>
                                    <option value="done">Done</option>
                                </select>
                                <select v-model="task.priority" class="w-full border p-2 rounded-lg mb-2">
                                    <option value="low">Low</option>
                                    <option value="medium">Medium</option>
                                    <option value="high">High</option>
                                </select>
                                <button @click="updateTask(task)"
                                    class="bg-green-500 text-white px-4 py-2 rounded-lg mr-2">Save</button>
                                <button @click="cancelEdit(task)"
                                    class="bg-red-500 text-white px-4 py-2 rounded-lg">Cancel</button>
                            </div>
                            <!-- View Mode -->
                            <div v-else class="flex justify-between">
                                <span>{{ task.name }}</span>
                                <div>
                                    <button @click="editTask(task)" class="text-blue-500 ml-2">Edit</button>
                                    <button @click="deleteTask(task)" class="text-red-500 ml-2">Delete</button>
                                </div>
                            </div>
                        </div>
                    </div>
                    <button @click="openAddTaskSidebar(column.name)"
                        class="bg-blue-500 text-white px-4 py-2 mt-4 w-full rounded-lg shadow-lg transition-all">
                        Add Task to {{ column.name }}
                    </button>
                </div>
            </div>
        </div>

        <!-- Add Task Sidebar -->
        <div v-if="showAddTaskSidebar"
            class="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center z-20">
            <div class="bg-white p-6 rounded-lg w-96">
                <h2 class="text-2xl font-semibold mb-4">Add Task to {{ currentStatus }}</h2>
                <input v-model="newTask.name" placeholder="Task Name" class="w-full border p-2 rounded-lg mb-4" />
                <textarea v-model="newTask.description" placeholder="Description"
                    class="w-full border p-2 rounded-lg mb-4"></textarea>
                <input v-model="newTask.dueDate" type="date" class="w-full border p-2 rounded-lg mb-4" />
                <select v-model="newTask.priority" class="w-full border p-2 rounded-lg mb-4">
                    <option value="low">Low</option>
                    <option value="medium">Medium</option>
                    <option value="high">High</option>
                </select>
                <button @click="addTask" class="bg-blue-500 text-white px-4 py-2 rounded-lg mr-2">Save Task</button>
                <button @click="toggleAddTaskSidebar" class="bg-red-500 text-white px-4 py-2 rounded-lg">Cancel</button>
            </div>
        </div>

        <!-- Dashboard with Pie Chart -->
        <div v-if="showDashboard" class="fixed inset-0 bg-black bg-opacity-50 flex justify-center items-center z-20">
            <div class="bg-white p-6 rounded-lg w-96">
                <h2 class="text-2xl font-semibold mb-4">Task Status</h2>
                <pie-chart :data="pieChartData"></pie-chart>
                <button @click="toggleDashboard" class="bg-red-500 text-white px-4 py-2 mt-4 w-full rounded-lg">
                    Close Dashboard
                </button>
            </div>
        </div>
    </div>
</template>

<script>
import axios from "axios";
import { Pie } from "vue-chartjs";
import { Chart as ChartJS } from "chart.js";
import { PieController, ArcElement, CategoryScale, Tooltip, Legend } from "chart.js";

ChartJS.register(PieController, ArcElement, CategoryScale, Tooltip, Legend);

export default {
    components: {
        "pie-chart": Pie,
    },
    data() {
        return {
            board: [
                { name: "To Do", tasks: [] },
                { name: "In Progress", tasks: [] },
                { name: "Done", tasks: [] },
            ],
            newTask: {
                name: "",
                description: "",
                dueDate: "",
                priority: "low",
            },
            currentStatus: "",
            showAddTaskSidebar: false,
            showDashboard: false,
            pieChartData: {
                labels: ["To Do", "In Progress", "Done"],
                datasets: [
                    {
                        data: [0, 0, 0],
                        backgroundColor: ["#ff6384", "#36a2eb", "#4bc0c0"],
                    },
                ],
            },
        };
    },
    methods: {
        async fetchTasks() {
            try {
                const response = await axios.get("http://localhost:8080/tasks");
                const tasks = response.data;

                tasks.forEach((task) => {
                    if (task.dueDate) {
                        task.dueDate = task.dueDate.split("T")[0]; // แปลง "2024-11-21T00:00:00Z" เป็น "2024-11-21"
                    }
                });

                this.board[0].tasks = tasks.filter((task) => task.status === "todo");
                this.board[1].tasks = tasks.filter((task) => task.status === "in-progress");
                this.board[2].tasks = tasks.filter((task) => task.status === "done");

                this.updatePieChartData();
            } catch (error) {
                console.error("Error fetching tasks:", error);
            }
        },

        logout() {
            localStorage.removeItem("token");
            this.$router.push("/login");
        },

        updatePieChartData() {
            this.pieChartData.datasets[0].data = [
                this.board[0].tasks.length,
                this.board[1].tasks.length,
                this.board[2].tasks.length,
            ];
        },
        toggleDashboard() {
            this.showDashboard = !this.showDashboard;
            if (this.showDashboard) {
                this.fetchTasks();
            }
        },
        openAddTaskSidebar(status) {
            this.currentStatus = status;
            this.showAddTaskSidebar = true;
        },
        toggleAddTaskSidebar() {
            this.showAddTaskSidebar = !this.showAddTaskSidebar;
        },
        async addTask() {
            const statusMap = {
                "To Do": "todo",
                "In Progress": "in-progress",
                "Done": "done",
            };

            const taskData = {
                name: this.newTask.name,
                status: statusMap[this.currentStatus],
                description: this.newTask.description,
                dueDate: this.newTask.dueDate,
                priority: this.newTask.priority,
            };

            try {
                await axios.post("http://localhost:8080/tasks", taskData);
                this.fetchTasks();
                this.toggleAddTaskSidebar();
                this.newTask = { name: "", description: "", dueDate: "", priority: "low" };
            } catch (error) {
                console.error("Error adding task:", error);
            }
        },
        editTask(task) {
            task.isEditing = true;
            task.originalData = { ...task }; // Backup ข้อมูลเดิม
        },
        async updateTask(task) {
            try {
                await axios.put(`http://localhost:8080/tasks/${task.id}`, task);
                task.isEditing = false;
                this.fetchTasks();
            } catch (error) {
                console.error("Error updating task:", error);
            }
        },
        cancelEdit(task) {
            Object.assign(task, task.originalData); // คืนค่าข้อมูลเดิม
            task.isEditing = false;
        },
        async deleteTask(task) {
            try {
                await axios.delete(`http://localhost:8080/tasks/${task.id}`);
                this.fetchTasks();
            } catch (error) {
                console.error("Error deleting task:", error);
            }
        },
    },
    mounted() {
        this.fetchTasks();
    },

};
</script>