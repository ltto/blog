// Generated by CoffeeScript 1.8.0
(function() {
  var set_comment_error;

  set_comment_error = function(error) {
    $('.comment_error').css('display', 'block');
    return $('.comment_error').html(error);
  };

  this.scroll_to_comment_form = function() {
    var comment_content_dom, scroll_to_top_position;
    comment_content_dom = $('#new_comment_form textarea');
    scroll_to_top_position = comment_content_dom.offset().top - 80;
    if (scroll_to_top_position < 0) {
      scroll_to_top_position = 0;
    }
    $('body,html').animate({
      scrollTop: scroll_to_top_position
    }, 800);
    return comment_content_dom.focus();
  };

  this.reply_comment = function(nickname, reply_to) {
    var content_dom, new_content, old_content, origin_comment_dom, to_append;
    content_dom = $('#new_comment_form textarea');
    old_content = content_dom.val() || '';
    if (nickname.indexOf(' ') === -1) {
      to_append = '@' + nickname + ' ';
    } else {
      to_append = '@' + nickname + ', ';
    }
    if (old_content.indexOf(to_append) !== -1) {
      to_append = '';
    }
    new_content = old_content + to_append;
    reply_to = reply_to || '';
    $('#reply_to_id').val(reply_to);
    if (reply_to) {
      origin_comment_dom = $('#' + reply_to);
      if (origin_comment_dom.length) {
        origin_comment_dom.append($('#new_comment_form'));
      }
    }
    content_dom.click();
    content_dom.val(new_content);
    scroll_to_comment_form();
    return false;
  };

  $(document).ready(function() {
    var author_dom, content_dom, email_dom, path_dom, site_dom;
    author_dom = $('#new_comment_form input[name="author"]');
    email_dom = $('#new_comment_form input[name="email"]');
    site_dom = $('#new_comment_form input[name="site"]');
    path_dom = $('#new_comment_form input[name="path"]');
    content_dom = $('#new_comment_form textarea');
    $(".new_comment").click(function() {
      $(".comment_trigger").hide();
      $(".comment_triggered").fadeIn("slow");
      if (author_dom.length) {
        author_dom.val(author_dom.val() || Cookies.get('comment_author') || '');
      }
      if (email_dom.length) {
        email_dom.val(email_dom.val() || Cookies.get('comment_email') || '');
      }
      if (site_dom.length) {
        return site_dom.val(site_dom.val() || Cookies.get('comment_site') || '');
      }
    });
    $('.new_comment textarea').keyup(function(event) {
      var current_height, new_comment_form;
      current_height = $(this).height();
      if (current_height < this.scrollHeight && current_height < 350) {
        $(this).height(this.scrollHeight);
      }
      if (event.which === 27) {
        new_comment_form = $('#new_comment_form');
        if (!new_comment_form.parent().hasClass('new_comment_form_container')) {
          $('.new_comment_form_container').append(new_comment_form);
          $('#reply_to_id').val('');
          return scroll_to_comment_form();
        }
      }
    });
    return $('.comment_submit_button').click(function() {
      var author, content, data_to_post, email, new_comment_form, parent_comment_id, site;
      author = author_dom.val() || '';
      email = email_dom.val() || '';
      site = site_dom.val() || '';
      content = content_dom.val();
      parent_comment_id = $('#reply_to_id').val() || '';
      new_comment_form = $('#new_comment_form');
      data_to_post = {
        author: author,
        email: email,
        site: site,
        content: content,
        path: path_dom.val(),
        reply: parent_comment_id,
        return_html: true
      };
      if (content.length < 5) {
        set_comment_error('min length of comment is 5!');
        content_dom.focus();
        return false;
      }
      if (!email && email_dom.length) {
        set_comment_error('email is required');
        content_dom.focus();
        return false;
      }
      if (author) {
        Cookies.set('comment_author', author, {
          expires: 9999
        });
      }
      if (email) {
        Cookies.set('comment_email', email, {
          expires: 9999
        });
      }
      if (site) {
        Cookies.set('comment_site', site, {
          expires: 9999
        });
      }
      $.ajax({
        url: new_comment_form.attr('action'),
        type: 'post',
        data: data_to_post,
        success: function(data) {
          var new_comment_dom, parent_comment_dom, sub_comments_dom, sub_comments_filter;
          $('.comment_error').css('display', 'none');
          if (data.error) {
            return set_comment_error(data.error);
          } else {
            if (typeof data === 'string') {
              content_dom.val('');
              if (parent_comment_id && $('#' + parent_comment_id).length) {
                parent_comment_dom = $('#' + parent_comment_id);
                sub_comments_filter = '#' + parent_comment_id + ' ul.sub_comments';
                if (!$(sub_comments_filter).length) {
                  parent_comment_dom.append('<ul class="sub_comments"></ul>');
                }
                sub_comments_dom = $(sub_comments_filter);
                sub_comments_dom.append(data);
                new_comment_dom = $(sub_comments_filter + ' .comment').last();
                $('.new_comment_form_container').append(new_comment_form);
              } else {
                $('.comments').append(data);
                new_comment_dom = $('.comments .comment').last();
              }
              $('html, body').animate({
                scrollTop: new_comment_dom.offset().top
              }, 500, 'swing', function() {
                return new_comment_dom.fadeIn(150).fadeOut(150).fadeIn(150);
              });
            }
            return console.log(data);
          }
        },
        fail: function(data) {
          console.log(data);
          return console.log('failed');
        }
      });
      return false;
    });
  });

}).call(this);
