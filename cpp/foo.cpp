#define LINUX_BUILD
#include <iostream>
#include <stdint.h>
#include <algorithm>    // std::sort
#include "DataDefs.h"
#include "Export.h"
#include "RemoteClient.h"
#include "RemoteTools.h"
#include "BasicApi.pb.h"
#include "MiscUtils.h"
#include "misc_trait_type.h"
#include "profession.h"
#include "job_type.h"
#include "job.h"
#include "job_item.h"
#include "unit_thought_type.h"
#include "unit_flags1.h"
#include "unit_flags2.h"
#include "modules/Materials.h"
#include "df/dfhack_material_category.h"
#include "foo.hpp"

//#include "library/include/df/item.h"
dfproto::ListUnitsIn *in;
dfproto::ListUnitsOut *out;
dfproto::GetWorldInfoOut *world_info_out;
dfproto::ListEnumsOut *enums_out;

bool myfunction (dfproto::BasicUnitInfo* dwarf1, dfproto::BasicUnitInfo* dwarf2 ){
    return dwarf1->profession() < dwarf2->profession(); 
}

std::vector<dfproto::BasicUnitInfo*> dwarves;
std::vector<CDwarf> cdwarves;
std::vector<std::string> elems = std::vector<std::string>();
DFHack::RemoteFunction<dfproto::ListUnitsIn, dfproto::ListUnitsOut> list_units;
DFHack::RemoteFunction<dfproto::EmptyMessage, dfproto::GetWorldInfoOut> get_world_info;
DFHack::color_ostream_wrapper * df_network_out = new DFHack::color_ostream_wrapper(std::cout);
DFHack::RemoteClient * network_client = new DFHack::RemoteClient(df_network_out);
std::string output;
std::string thought_string = "";

void cxxFoo::Init(void) {
    network_client->connect();
    DFHack::RemoteFunction<dfproto::EmptyMessage, dfproto::ListEnumsOut> list_enums;
    list_units.bind(network_client, "ListUnits");
    get_world_info.bind(network_client, "GetWorldInfo");
    list_enums.bind(network_client, "ListEnums");
    in = new dfproto::ListUnitsIn();
    out = new dfproto::ListUnitsOut();
    in->set_scan_all(true);
    in->set_race(465);
    in->set_alive(true);
    in->mutable_mask()->set_profession(true);
    in->mutable_mask()->set_misc_traits(true);
    //my_call(network_client->default_output(), dfproto::ListUnitsIn::default_instance, dfproto::ListUnitsOut::default_instance);
    world_info_out = new dfproto::GetWorldInfoOut();

    std::ostringstream stream;
    DFHack::color_ostream_wrapper *df_output = new DFHack::color_ostream_wrapper(stream);
    network_client->run_command((*df_output), "collect_reactions", std::vector<std::string>());


    std::string item;
    std::istringstream iss(stream.str());
    while (std::getline(iss, item, '\n')) {
        elems.push_back(item);
    }
    std::cout << std::to_string(elems.size()) + "\n";

    enums_out = new dfproto::ListEnumsOut();
    list_enums(new dfproto::EmptyMessage(), enums_out);

	std::cout<<this->a<<std::endl;
}

CDwarf cxxFoo::GetDwarf(int i){
    return cdwarves.at(i);
}

bool cxxFoo::Paused() {
    return world_info_out->pause_state();
}

void cxxFoo::Exit() {
    network_client->disconnect();
}

// TODO: size of elems can change based on world
const char* cxxFoo::GetJobType(int i) {
    /*std::string job_type = enums_out->job_type_type(i).name();
    std::string item_type = enums_out->job_type_material(i).name();
    if (job_type == "8"){//
        return(enums_out->job_type_caption(i).name() + " " + item_type).c_str();
    } else {
        return " ";
    }*/
    if(i < elems.size()){
        return elems[i].c_str();
    } else {
        return " ";
    }
}

const char* cxxFoo::GetFirstName(int i) {
    dfproto::BasicUnitInfo* dwarf = dwarves.at(i);
    return output.c_str();
    //return (dwarf->name().first_name() + " " + dwarf->name().last_name() + " " + dwarf->thought_string()).c_str();
}

int cxxFoo::GetId(int i) {
    dfproto::BasicUnitInfo* dwarf = dwarves.at(i);
    return dwarf->unit_id();
}

const char* DwarfGetHappiness(dfproto::BasicUnitInfo* dwarf){
    int happy = dwarf->happiness();
    int fg;
    const char* level;
    if (happy == 0)         
        level = "miserable";
    else if (happy <= 25)   // 
        level = "very unhappy";
    else if (happy <= 50)   // 
        level = "unhappy";
    else if (happy < 75)    // 
        level = "fine";
    else if (happy < 125)   // 
        level = "quite content";
    else if (happy < 150)   // 
        level = "happy";
    else                    // ecstatic
        level = "escstatic";

    return level;
}

const char* cxxFoo::GetHappiness(int i) {
    dfproto::BasicUnitInfo* dwarf = dwarves.at(i);
    return DwarfGetHappiness(dwarf);
}

const char* cxxFoo::GetCurrentJob(int i) {
    dfproto::BasicUnitInfo* dwarf = dwarves.at(i);
    for(int i = 0; i < dwarf->misc_traits_size(); i++){
        switch(dwarf->misc_traits(i).id()){
            case df::misc_trait_type::OnBreak:
                return "On Break";
            case df::misc_trait_type::Migrant:
                return "New Arrival";
        }
    }
    output = dwarf->current_job();
    return output.c_str();
}

int cxxFoo::Size() {
    return out->value_size();
}

const char* DwarfGender(dfproto::BasicUnitInfo* dwarf){
    if(dwarf->gender() == 1){
        return "He";
    } else {
        return "She";
    }
}

const char* cxxFoo::Gender(int i) {
    dfproto::BasicUnitInfo* dwarf = dwarves.at(i);
    return DwarfGender(dwarf);
}


const char* DwarfGetThoughts(dfproto::BasicUnitInfo* dwarf) {
    thought_string = DwarfGender(dwarf);
    thought_string += " has been "; 
    thought_string += DwarfGetHappiness(dwarf);
    thought_string += " lately. \n";
    for(int i = 0; i < dwarf->recent_events_size(); i++){
        const dfproto::UnitThought* thought = &dwarf->recent_events(i);
        thought_string += DwarfGender(dwarf);
        thought_string += " ";
        thought_string += enums_out->unit_thought_type_caption(thought->type()).name(); // << std::endl;
        if(thought->type() == df::unit_thought_type::Talked){
            std::string relationship = "a " + enums_out->unit_relationship_type(thought->subtype()).name();
            thought_string.replace(thought_string.find("(somebody/a pet/a spouse/...)"), sizeof("(somebody/a pet/a spouse/...)")-1, relationship);
        } else if(thought->type() == df::unit_thought_type::AdmireBuilding){
            std::string level = "fine";
            std::string building = enums_out->building_type(thought->subtype()).name();
            thought_string.replace(thought_string.find("(building)"), sizeof("(building)")-1, building);
            thought_string.replace(thought_string.find("(fine/very fine/splendid/wonderful/completely sublime)"), sizeof("(fine/very fine/splendid/wonderful/completely sublime)")-1, level);
        }
            //thought_string += ENUM_ATTR_STR(unit_thought_type, caption, thought->type);
            //thought_string += std::to_string(thought->subtype);
            //thought_string += std::to_string(thought->severity);
        //thought_string += std::to_string(thought->age());
        thought_string += ".\n";
    }
    return thought_string.c_str();
}

const char* cxxFoo::GetThoughts(int j) {
    dfproto::BasicUnitInfo* dwarf = dwarves.at(j);
    return DwarfGetThoughts(dwarf);
}

void cxxFoo::Update() {
    //out = new dfproto::ListUnitsOut();
    int result = get_world_info(new dfproto::EmptyMessage(), world_info_out);
    if(result != 0){
        std::cout << " failed to get world!" << result << "\n";
        network_client->disconnect();
        exit(1);
    }
    list_units(in, out);
    dwarves.clear();
    cdwarves.clear();
    dfproto::BasicUnitInfo* dwarf;
    for(int i = 0; i < out->value_size(); i++){
        dwarf = (dfproto::BasicUnitInfo*)&out->value(i);
        df::unit_flags1 flags = dwarf->flags1();
        df::unit_flags2 flags2 = dwarf->flags2();
        if (!flags.bits.merchant && !flags.bits.diplomat && !flags2.bits.visitor){
            dwarves.push_back(dwarf);
            output = dwarf->name().first_name() + " " + dwarf->name().last_name() + " \n" + dwarf->profession_name(); 
            output[0] = toupper(output[0]);
            CDwarf cdwarf = { 
                .name = output.c_str(),
                .thoughts = DwarfGetThoughts(dwarf)
            };
            cdwarves.push_back(cdwarf);
        } else {
            std::cout << "merchant\n";
        }
    }
    std::sort (dwarves.begin(), dwarves.end(), myfunction);
}

