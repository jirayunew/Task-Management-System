<template>
    <div class="h-screen flex justify-center items-center bg-gray-100">
        <div class="bg-white p-6 rounded-lg shadow-xl w-96">
            <h2 class="text-2xl font-bold mb-4 text-center p-2">Login</h2>
            <input v-model="username" placeholder="Username" class="w-full p-2 mb-4 border rounded-lg" />
            <input v-model="password" type="password" placeholder="Password"
                class="w-full p-2 mb-4 border rounded-lg" />
            <button @click="login" class="bg-blue-500 text-white w-full p-2 rounded-lg">Login</button>
            <p v-if="error" class="text-red-500 mt-2">{{ error }}</p>
        </div>
    </div>
</template>

<script>
import axios from "axios";

export default {
    data() {
        return {
            username: "",
            password: "",
            error: null,
        };
    },
    methods: {
        async login() {
            try {
                const response = await axios.post("http://localhost:8080/login", {
                    username: this.username,
                    password: this.password,
                });
                localStorage.setItem("token", response.data.token); // เก็บ JWT ลงใน localStorage
                this.$router.push("/"); // Redirect ไปหน้า App
            } catch (err) {
                this.error = "Invalid username or password";
            }
        }
    },
};
</script>

<style scoped>
/* Add styles here */
</style>