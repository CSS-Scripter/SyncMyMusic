<template>
    <header>
        <router-link to="/"><h1>SyncMyMusic</h1></router-link>
        <router-link class="nav-button" to="/account" v-if=isLoggedIn>Account</router-link>
        <router-link class="nav-button" to="/login" v-else id="login-button">Log in</router-link>
    </header>
</template>

<script>
export default {
    computed: {
        isLoggedIn: function() {
            return this.$store.getters.isLoggedIn;
        }
    },
}
</script>

<style lang="scss" scoped>
    @import '../variables';
    template {
        width: 100%;
    }
    header {
        width: 100%;
        max-width: 1000px;
        margin: 0 auto;
        font-family: 'Rubik', sans-serif;
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        padding: 1em 1.2em;

        h1 {
            font-family: 'Saira', sans-serif;
            font-weight: bold;
            font-style: normal;
            font-size: 36px;
            color: $light_color;
            margin: 0;
            padding: 0;
        }

        a {
            color: $light_color;
            font-weight: bold;
            text-decoration: none;
        
            &.nav-button {
                font-size: 24px;
                margin: 14px 0;
            
                &:hover {
                    text-decoration: underline;
                }
            }

            &:hover {
                cursor: pointer;
            }
        }
    }
</style>