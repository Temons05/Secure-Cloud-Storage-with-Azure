<template>
  <div class="login-page d-grid">
    <b-container fluid class="login-container">
      <b-row class="h-100">
        <b-col :order="isLogin ? 0 : 1" class="left-container">
          <b-row align-v="center" class="h-100 left-container-inner">
            <b-col>
              <b-row>
                <h1 class="login_title">
                  {{ isLogin ? "Sign in" : "Register" }}
                </h1>
              </b-row>
              <b-row v-if="authStatus">
                <h1 class="error-message">
                  {{ authStatus }}
                </h1>
              </b-row>
              <b-row>
                <b-input
                  class="login-input"
                  aria-label="Username"
                  placeholder="Username"
                  type="text"
                  v-model="username"
                />
              </b-row>
              <b-row>
                <b-input
                  class="login-input"
                  aria-label="password"
                  placeholder="Password"
                  type="password"
                  v-model="password"
                />
              </b-row>
              <b-row>
                <button
                  @click="isLogin ? login() : register()"
                  class="login_btn mx-auto"
                >
                  {{ isLogin ? "Sign in" : "Register" }}
                </button>
              </b-row>
            </b-col>
          </b-row>
        </b-col>

        <b-col
          class="right-container h-100"
          :style="{ borderRadius: isLogin ? '0 26px 26px 0' : '26px 0 0 26px' }"
        >
          <b-row align-v="center" class="right-container-inner h-100">
            <b-col>
              <b-row>
                <h1 class="login_title">
                  {{
                    isLogin
                      ? "Don't have an account?"
                      : "Already have an account?"
                  }}
                </h1>
              </b-row>
              <b-row>
                <p class="login_desc">
                  {{ isLogin ? "Register now!" : "Login now!" }}
                </p>
              </b-row>
              <b-row>
                <button @click="toggleLoginSignup" class="register_btn mx-auto">
                  {{ isLogin ? "Sign up" : "Log in" }}
                </button>
              </b-row>
            </b-col>
          </b-row>
        </b-col>
      </b-row>
    </b-container>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { Component } from "vue-property-decorator";
import Navbar from "@/components/Navbar.vue";
import { post } from "@/utils/utils";

@Component({
  name: "Home",
  components: {
    Navbar,
  },
})
export default class Home extends Vue {
  private isLogin = true;

  private username = "";
  private password = "";

  private authStatus = "";

  private created() {
    if (this.$cookies.get("jwtToken")) {
      this.$router.push("/dashboard");
    }
  }

  private toggleLoginSignup() {
    this.isLogin = !this.isLogin;
  }

  private async login() {
    try {
      const resp = (await (
        await post("/login", {
          username: this.username,
          password: this.password,
        })
      ).json()) as LoginResponse;

      this.$cookies.set("jwtToken", resp.data.token, 3 * 60 * 60);
      console.log(resp);

      this.$router.push("/dashboard");
    } catch (e) {
      console.error(e);
      this.authStatus = "Invalid Username/Password";
    }
  }

  private async register() {
    const resp = (await (
      await post("/register", {
        username: this.username,
        password: this.password,
      })
    ).json()) as RegisterResponse;

    console.log(resp);

    if (resp.error) {
      this.authStatus = resp.error;
    } else if (resp.success) {
      this.authStatus = resp.success
        ? "Successfully registered"
        : "Registeration failed";
      this.toggleLoginSignup();
    }
  }
}
</script>

<style lang="sass" scoped>
.login-container
  box-shadow: 4px 7px 12px -10px rgba(0, 0, 0, 0.25), 0 10px 10px rgba(0, 0, 0, 0.22)
  text-align: center
  border-radius: 26px
  background-color: #ffffff
  width: 850px
  height: 600px

.right-container
  color: white
  text-align: center
  background: #834ce6
  background-color: #3787f6
  background-image: linear-gradient(315deg, #733fed 0%, #3787f6 74%)
  background-repeat: no-repeat
  background-size: cover
  background-position: 0 0

.left-container
  border-radius: 26px

.login-page
  background: #f6f5f7
  height: 100vh
  place-items: center

.left-container-inner, .right-container-inner
  padding: 35px

.login-input
  background-color: #eee !important
  border: none
  padding: 12px 15px
  margin: 20px 0 0 0
  width: 100%
  border-radius: 5px

.forgot_pass
  color: #333
  font-size: 14px
  text-decoration: none
  margin: 20px 0

.login_btn
  border: 1px solid #7853ce
  background-color: #735cdd
  margin-top: 25px

.login_btn, .register_btn
  border-radius: 20px
  width: calc(100% - 130px)
  color: #ffffff
  font-size: 12px
  font-weight: bold
  padding: 12px 45px
  letter-spacing: 1px
  text-transform: uppercase

.register_btn
  border: 1px solid #fff
  background-color: transparent

.login_title
  margin-bottom: 25px
  font-weight: bold

.login_desc
  font-size: 15px
  font-weight: 400
  line-height: 20px
  letter-spacing: 0.5px
  margin: 20px 0 30px

.error-message
  color: crimson
  margin-bottom: 20px
  font-size: 14px
</style>
