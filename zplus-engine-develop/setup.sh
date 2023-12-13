#! bash
PROJECT_ROOT_DIRECTORY=$(pwd)

TOOL_NAME_JMETER=JMeterLoadTesting
TOOL_NAME_METASPLOIT=Metasploit
TOOL_NAME_NMAP=Nmap
TOOL_NAME_OPENVAS=OpenVAS
TOOL_NAME_OWASPDEPENDENCY=OWASPDependencyCheck
TOOL_NAME_OWASPZAP=OWASPZAP
TOOL_NAME_SQLMAP=SQLMap

SECURITY_TOOL_VOLUME_ROOT_DIRECTORY=$PROJECT_ROOT_DIRECTORY/security_tool_volume
SECURITY_TOOL_ENGINE_INPUT_FOLDER=$SECURITY_TOOL_VOLUME_ROOT_DIRECTORY/inputs
SECURITY_TOOL_ENGINE_HELPER_FOLDER=$SECURITY_TOOL_VOLUME_ROOT_DIRECTORY/helper
SECURITY_TOOL_ENGINE_RESULTS_FOLDER=$SECURITY_TOOL_VOLUME_ROOT_DIRECTORY/results

SECURITY_TOOL_ENGINE_OPENVAS_SOCKET_FOLDER=$SECURITY_TOOL_VOLUME_ROOT_DIRECTORY/openvas_socket
SECURITY_TOOL_ENGINE_OWASP_ZAP_FOLDER=$SECURITY_TOOL_VOLUME_ROOT_DIRECTORY/$TOOL_NAME_OWASPZAP

SECURITY_TOOL_ENGINE_NMAP_RESULTS_FOLDER=$SECURITY_TOOL_ENGINE_RESULTS_FOLDER/$TOOL_NAME_NMAP
SECURITY_TOOL_ENGINE_OPENVAS_RESULTS_FOLDER=$SECURITY_TOOL_ENGINE_RESULTS_FOLDER/$TOOL_NAME_OPENVAS
SECURITY_TOOL_ENGINE_OPENVAS_RESULTS_FOLDER=$SECURITY_TOOL_ENGINE_RESULTS_FOLDER/$TOOL_NAME_METASPLOIT

SECURITY_TOOL_ENGINE_OWASP_DEPENDENCY_CHECK_RESULTS_FOLDER=$SECURITY_TOOL_ENGINE_RESULTS_FOLDER/$TOOL_NAME_OWASPDEPENDENCY

SECURITY_TOOL_ENGINE_SQLMAP_RESULTS_FOLDER=$SECURITY_TOOL_ENGINE_RESULTS_FOLDER/$TOOL_NAME_SQLMAP

SECURITY_TOOL_ENGINE_JMETER_INPUTS_FOLDER=$SECURITY_TOOL_ENGINE_INPUT_FOLDER/$TOOL_NAME_JMETER
SECURITY_TOOL_ENGINE_JMETER_LOGS_FOLDER=$SECURITY_TOOL_ENGINE_RESULTS_FOLDER/$TOOL_NAME_JMETER/logs
SECURITY_TOOL_ENGINE_JMETER_REPORTS_FOLDER=$SECURITY_TOOL_ENGINE_RESULTS_FOLDER/$TOOL_NAME_JMETER/reports

clean_all=$1

if [[ $# -gt 0 && ($clean_all == "clean-all") ]]
then
    rm -rf $SECURITY_TOOL_VOLUME_ROOT_DIRECTORY
fi

mkdir -p $SECURITY_TOOL_ENGINE_INPUT_FOLDER
mkdir -p $SECURITY_TOOL_VOLUME_ROOT_DIRECTORY
mkdir -p $SECURITY_TOOL_ENGINE_HELPER_FOLDER
mkdir -p $SECURITY_TOOL_ENGINE_RESULTS_FOLDER

mkdir -p $SECURITY_TOOL_ENGINE_OPENVAS_SOCKET_FOLDER
mkdir -p $SECURITY_TOOL_ENGINE_OWASP_ZAP_FOLDER

mkdir -p $SECURITY_TOOL_ENGINE_NMAP_RESULTS_FOLDER
mkdir -p $SECURITY_TOOL_ENGINE_SQLMAP_RESULTS_FOLDER
mkdir -p $SECURITY_TOOL_ENGINE_OPENVAS_RESULTS_FOLDER

mkdir -p $SECURITY_TOOL_ENGINE_OWASP_DEPENDENCY_CHECK_RESULTS_FOLDER
mkdir -p $SECURITY_TOOL_ENGINE_JMETER_INPUTS_FOLDER
mkdir -p $SECURITY_TOOL_ENGINE_JMETER_LOGS_FOLDER
mkdir -p $SECURITY_TOOL_ENGINE_JMETER_REPORTS_FOLDER

sudo chmod -R 777 $SECURITY_TOOL_VOLUME_ROOT_DIRECTORY
