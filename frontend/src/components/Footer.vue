<template>
    <footer>
        <div class="content">
            <div class="links">
                <router-link to="/">About us</router-link>
                <router-link to="/">Terms of Use</router-link>
                <router-link to="/">Privacy</router-link>
            </div>
            <div class="copyright">Copyright C 2020 SyncMyMusic, All Rights reserved</div>
        </div>
    </footer>
</template>

<style lang="scss" scoped>
    footer {
        width: 100%;
        background-color: #002;
        padding: 1em 0 .8em 0;
    }

    .content {
        display: flex;
        flex-direction: row;
        justify-content: space-between;

        width: 100%;
        max-width: 1000px;
        margin: 0 auto;

        color: #DDF;
        font-weight: lighter;
        font-size: 14px;
        font-family: 'Rubik', sans-serif;
    
        .links {
            display: grid;
            grid-template-columns: repeat(3, 1fr);
            grid-column-gap: 2em;

            a {
                color: #DDF;
                text-align: center;
            }
        }
        .copyright {
            padding-right: 1em;
        }
    }
</style>