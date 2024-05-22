# Changelog

## v1.3.1 (2024-05-22)

- add Marshal functions for struct with many types
  - ChatBoostSource
  - ChatBackground
  - ChatMember
  - MenuButton
  - MaybeInaccessibleMessage
  - ReactionType
  - MessageOrigin

## v1.3.0 (2024-05-20)

- support API v7.3

## v1.2.2 (2024-04-25)

- fix race in test 
- add example inline_keyboard_multiselect 

## v1.2.1 (2024-04-03)

- Added error handling for specific error codes (#69) 
- fix: exit from waitUpdates function when context is terminated (#75) 

## v1.2.0 (2024-04-02)

- support API v7.2
- remove worker pool
- add option `WithUpdatesChannelCap`

## v1.1.7 (2024-03-11)

- add error `ErrorForbidden` for responses with error code 403
- fix CopyMessages capture response 
- ForwardMessages properly unmarshal result (#68) 

## v1.1.6 (2024-03-05)

- ForwardMessages return slice of MessageID (#66)

## v1.1.5 (2024-03-04)

- add option WithAllowedUpdates (#65)

## v1.1.4 (2024-02-27)

- fix `CallbackQuery.Message` field type from `InaccessibleMessage` to `MaybeInaccessibleMessage`

## v1.1.3 (2024-02-22)

- fix json tag for DeleteMessagesParams.MessageIDs

## v1.1.2 (2024-02-21)

- fix SetMessageReaction marshal func, fix get response for this method

## v1.1.1 (2024-02-21)

- fix typo in json tag `models.UsersShared` for field UserIDs 

## v1.1.0 (2024-02-19)

- support API v7.1

## v1.0.1 (2024-01-15)

- fix: add 'omitempty' to LinkPreviewOptions.URL

## v1.0.0 (2024-01-10)

- support API v7.0

## v0.8.3 (2023-12-07)

- Fix typo for setChatDescription params (#49)
- add example handler_match_func
- add function bot.FileDownloadLink

## v0.8.2 (2023-10-16)

- add missed field HasSpoiler to InputMedia types

## v0.8.1 (2023-10-09)

- json tag typo fix (#41) vvok12* 

## v0.8.0 (2023-09-28)

- support bot api 6.9 

## v0.7.15 (2023-08-24)

- support bot api 6.8 

## v0.7.14 (2023-08-14)

- update `defaultUpdatesChanCap` from 64 to 1024

## v0.7.13 (2023-06-14)

- revert `v0.7.7` changes

## v0.7.12 (2023-06-13)

- check for empty token on bot init

## v0.7.11 (2023-05-26)

- add missed ChatID to BotCommandScopeChat
- add missed BotCommandScopeChatAdministrators.MarshalCustom

## v0.7.10 (2023-05-22)

- add ChatAction type with consts for `sendChatAction` method

## v0.7.9 (2023-05-16)

- add option `WithSkipGetMe()`

## v0.7.8 (2023-04-26)

- [BUGFIX] fix type of field `Message.ReplyMarkup` (#30)
- improve examples (#31)

## v0.7.7 (2023-04-24)

- change `models.ReplyMarkup` from `any` to interface

## v0.7.6 (2023-04-24)

- support bot api 6.7

## v0.7.5 (2023-04-13)

- add option `WithDebugHandler`

## v0.7.4 (2023-04-10)

- [BUGFIX] change field `SendPollParams.IsAnonymous` to *bool (#26)
- add helpers `bot.True()` and `bot.False()` for define *bool values

## v0.7.3 (2023-04-05)

- make `bot.ProcessUpdate` public 

## v0.7.2 (2023-04-03)

- fix copyMessage method (#24)

## v0.7.1 (2023-03-20)

- fix field name in MenuButtonWebApp #21

## v0.7.0 (2023-03-14)

- support bot api 6.6

## v0.6.0 (2023-03-01)

- [BREAKING] change UserID field type in methods params and models from `int` to `int64` #19

## v0.5.1 (2023-02-10)

- [BUGFIX] send http.NoBody if all params fields are empty. For example, method `getMyCommands` has not required fields

## v0.5.0 (2023-02-06)

- support bot api 6.5 

## v0.4.3 (2023-02-05)

- [BUGFIX] SendChatAction now calls `sendChatAction` instead `sendDice`

## v0.4.2 (2023-01-30)

- [BUGFIX] change field Thumb to InputFile type

## v0.4.1 (2023-01-19)

- add func `EscapeMarkdownUnescaped` for escape only unescaped symbols

## v0.4.0 (2023-01-13)

- [BREAKING] change some types from `int` to `int64` #11

## v0.3.4 (2023-01-10)

- support bot api 6.4

## v0.3.3 (2022-12-28)

- add `RegisterHandlerMatchFunc` 

## v0.3.2 (2022-11-28)

- fix field type SendDocumentParams.Document from string to InputFile, add example

## v0.3.1 (2022-11-05)

- support bot api 6.3

## v0.3.0 (2022-10-14)

- **[API CHANGE]** method New now call `getMe` after init and returns error as second result on fail

## v0.2.2 (2022-08-23)

- support bot api v6.2

## v0.2.1 (2022-07-08)

- support bot api v6.1

## v0.2.0 (2022-05-06)

- change API. Use bot funcs instead of`methods` package.

## v0.1.1 (2022-05-02)

- add webhook support
- API refactoring
  - use bot.GetUpdates() instead of bot.Start()

## v0.1.0 (2022-05-02)

- initial release
