export INTERVIEW_KEY=$(grep "interview-"/mystery/*)
echo $INTERVIEW_KEY
cat /mystery/interview_$INTERVIEW_KEY
echo "$MAIN_SUSPECT"