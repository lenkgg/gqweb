<script>
var api = require('../utils/api');
import {CreateScopes} from '../utils/hedge/scopes'
import {CreateSymlist} from '../utils/symbol/symlist'
import {CreateUser} from '../utils/userInfo'
import {CreateMystocks} from '../utils/symbol/mystocks'

let defaultIndexs = CreateSymlist("index")
let userInfo = CreateUser()
let hedgeScopes = CreateScopes()
let mystocks = CreateMystocks()

export default {
  userInfo,
  url: 'https://www.lilishare.com:8443/',
  hedgeScopes,
  defaultIndexs,
  mystocks,
}

</script>
