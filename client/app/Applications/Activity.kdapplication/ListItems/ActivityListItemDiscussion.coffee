class DiscussionActivityItemView extends ActivityItemChild

  constructor:(options, data)->

    options = $.extend
      cssClass    : "activity-item discussion"
      tooltip     :
        title     : "Discussion"
        offset    : 3
        selector  : "span.type-icon"
    ,options

    super options,data

    @actionLinks = new DiscussionActivityActionsView
      delegate : @commentBox.opinionList
      cssClass : "reply-header"
    , data

    # the ReplyIsAdded event is emitted by the JDiscussion model in bongo
    # with the object references to author/origin and so on in the reply
    # argument. if the new reply is supposed to be added to the client data
    # structure, then it must be created as a new  JOpinion and then populated
    # by the data on the reply.opinionData field (JSON of the actual object)

    data.on 'ReplyIsAdded', (reply)=>

      if data.bongo_.constructorName is "JDiscussion"

        # This would add the actual items to the views once posted
        #
        # Why this workaround, you ask?
        #
        #  Getting the data from the JDiscussion.reply event "ReplyIsAdded"
        #  without JSONifying it locks up the UI for up to 10 seconds.

        # Create new JOpinion and convert JSON into Object

        # newOpinion = new bongo.api.JOpinion
        # opinionData = JSON.parse(reply.opinionData)

        # Copy JSON data to the newly created JOpinion

        # for variable of opinionData
        #   newOpinion[variable] = opinionData[variable]

        # Updating the local data object, then adding the item to the box
        # and increasing the count box

        # if data.opinions?
        #   unless data.opinions.indexOf newOpinion is -1
        #     data.opinions.push newOpinion
        # else
        #   data.opinions = [newOpinion]

        # The following line would add the new Opinion to the View
        # @opinionBox.opinionList.addItem newOpinion, null, {type : "slideDown", duration : 100}

        # unless reply.replier.id is KD.whoami().getId()
        @opinionBox.opinionList.emit "NewOpinionHasArrived"

    @opinionBox = new DiscussionActivityOpinionView
      cssClass    : "activity-opinion-list comment-container"
    , data

    # When an opinion gets deleted, then the removeReply method of JDiscussion
    # will emit this event. This is a workaround for the OpinionIsDeleted
    # event not being caught for opinions that are loaded to the client data
    # structure after the snapshot is loaded

    data.on "ReplyIsRemoved",(replyId)=>

      # this will remove the item from the list if the data doesn't
      # contain it anymore, but the list does. the next snapshot refresh
      # will be okay
      # This is needed, because the "OpinionIsDeleted" event isn't available
      # for newly added JOpinions, for some reason. --arvid

      for item,i in @opinionBox.opinionList.items
        if item?.getData()._id is replyId
          item.hide()
          item.destroy()

  viewAppended:()->
    return if @getData().constructor is KD.remote.api.CDiscussionActivity
    super()
    @setTemplate @pistachio()
    @template.update()

    # Here, the maxheight-reliant "View full discussion"-bar is toggled.
    # The shortened text is not sufficient since it can contain 500 line breaks
    # or <code> with very high whitespace amount. This keeps the snapshot view
    # clean.

    if @$("p.comment-body").height() >= 250
      @$("div.view-full-discussion").show()
    else
      @$("div.view-full-discussion").hide()

  render:->
    super()

  click:(event)->
    if $(event.target).closest("[data-paths~=title],[data-paths~=body]")
      if not $(event.target).is("a.action-link, a.count, .like-view")
        KD.getSingleton('router').handleRoute "/Activity/#{@getData().slug}", state:@getData()
        #appManager.tell "Activity", "createContentDisplay", @getData()

  applyTextExpansions:(str = "")->
    str = @utils.expandUsernames str

    if str.length > 500
      visiblePart = str.substr 0, 500
      # this breaks the markdown sanitizer
      # morePart = "<span class='more'><a href='#' class='more-link'>show more...</a>#{str.substr 501}<a href='#' class='less-link'>...show less</a></span>"
      str = visiblePart  + " ..." #+ morePart

    return str

  pistachio:->
    """
    <div class="activity-discussion-container">
      <span class="avatar">{{> @avatar}}</span>
      <div class='activity-item-right-col'>
        {{> @settingsButton}}
        <h3 class='hidden'></h3>
        <p class="comment-title">{{@applyTextExpansions #(title)}}</p>
        <footer class='clearfix'>
          <div class='type-and-time'>
            <span class='type-icon'></span> by {{> @author}}
            <time>{{$.timeago #(meta.createdAt)}}</time>
            {{> @tags}}
          </div>
          {{> @actionLinks}}
        </footer>
        {{> @opinionBox}}
      </div>
    </div>
    """

