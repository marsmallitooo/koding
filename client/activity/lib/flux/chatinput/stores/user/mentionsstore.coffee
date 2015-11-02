KodingFluxStore = require 'app/flux/base/store'
toImmutable     = require 'app/util/toImmutable'

###*
 * Store to handle user mentions
###
module.exports = class ChatInputUserMentionsStore extends KodingFluxStore

  @getterPath = 'ChatInputUserMentionsStore'


  getInitialState: ->

    mentions = [
      {
        names       : [ 'channel', 'all', 'team', 'group' ]
        description : 'notify everyone in the channel'
      }
      {
        names       : [ 'admins' ]
        description : 'notify only team admins'
      }
    ]

    return toImmutable mentions

