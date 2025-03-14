package promptservice_test

//func TestAddPromptCreatesThread(t *testing.T) {
//	univ, err := di.BigBang(&di.Options{
//		Test: true,
//	})
//	require.NoError(t, err)
//	ps := univ.PromptService
//
//	promptId := uuid.New().String()
//	threadId := uuid.New().String()
//	t.Run("add prompt", func(t *testing.T) {
//		_, err := ps.AddPrompt(context.Background(), &prompttypes.AddPromptRequest{
//			PromptCreateFields: prompttypes.PromptCreateFields{
//				Id:       promptId,
//				Prompt:   "hi",
//				ThreadId: threadId,
//			},
//		}, "")
//		require.NoError(t, err)
//	})
//}
